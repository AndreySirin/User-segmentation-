package storage

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/AndreySirin/avito-backend-assignment-2023/internal/entity"
	"github.com/AndreySirin/avito-backend-assignment-2023/internal/logger"
)

type SubscriptionStorage struct {
	lg *logger.MyLogger
	db *sql.DB
}

func NewSubscription(db *Storage) *SubscriptionStorage {
	return &SubscriptionStorage{
		lg: db.lg,
		db: db.db,
	}
}

func (s *SubscriptionStorage) InsertUserInSegment(ctx context.Context, subs entity.CreateSubscription) (err error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("error starting transaction: %w", err)
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	var exists bool
	err = tx.QueryRowContext(ctx, "SELECT EXISTS(SELECT 1 FROM users WHERE id_user = $1)", subs.IdUser).
		Scan(&exists)
	if err != nil {
		return fmt.Errorf("error checking if user exists: %w", err)
	}
	if !exists {
		return fmt.Errorf("user does not exist")
	}

	rows, err := tx.QueryContext(
		ctx,
		"SELECT id_segment FROM segments WHERE title=ANY($1)",
		subs.NameSegment,
	)
	if err != nil {
		return fmt.Errorf("error selecting segment: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var i int
		err = rows.Scan(&i)
		if err != nil {
			return fmt.Errorf("error scanning segment: %w", err)
		}
		subs.IdSegment = append(subs.IdSegment, i)
	}
	if err = rows.Err(); err != nil {
		return fmt.Errorf("failed to iterate rows: %w", err)
	}

	_, err = tx.ExecContext(ctx, `
INSERT INTO subscriptions (id_user,id_segment)
SELECT $1, id_segm
FROM UNNEST($2::int[]) AS id_segm`,
		subs.IdUser, subs.IdSegment)
	if err != nil {
		return fmt.Errorf("error inserting segment: %w", err)
	}
	return nil
}

func (s *SubscriptionStorage) DeleteUserInSegment(ctx context.Context, subs entity.CreateSubscription) (err error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("error starting transaction: %w", err)
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	rows, err := tx.QueryContext(
		ctx,
		"SELECT id_segment FROM segments WHERE title=ANY($1)",
		subs.NameSegment,
	)
	if err != nil {
		return fmt.Errorf("error selecting segment: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var i int
		err = rows.Scan(&i)
		if err != nil {
			return fmt.Errorf("error scanning segment: %w", err)
		}
		subs.IdSegment = append(subs.IdSegment, i)
	}
	if err = rows.Err(); err != nil {
		return fmt.Errorf("failed to iterate rows: %w", err)
	}
	res, err := tx.ExecContext(
		ctx,
		`DELETE FROM subscriptions WHERE id_segment=ANY($1)`,
		subs.IdSegment,
	)
	if err != nil {
		return fmt.Errorf("error deleting segment: %w", err)
	}
	row, err := res.RowsAffected()
	if err != nil || row == 0 {
		return fmt.Errorf("error getting rows affected: %w", err)
	}
	return nil
}
