package storage

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/AndreySirin/avito-backend-assignment-2023/internal/entity"
)

func (s *Storage) InsertSubscription(
	ctx context.Context,
	tx *sql.Tx,
	subs *entity.Subscription,
) error {
	_, err := tx.ExecContext(ctx, `
INSERT INTO subscriptions (user_id,segment_id,ttl,is_auto_add)
SELECT $1, segment_id,ttl,is_auto_add
FROM UNNEST($2::uuid[], $3::timestamp[], $4::bool[]) AS u(segment_id, ttl, is_auto_add)`,
		subs.IDUser, subs.IDSegment, subs.TTL, subs.AutoAdded)
	if err != nil {
		return fmt.Errorf("error inserting segment: %w", err)
	}
	return nil
}

func (s *Storage) DeleteSubscription(
	ctx context.Context,
	tx *sql.Tx,
	subs *entity.Subscription,
) error {
	res, err := tx.ExecContext(
		ctx,
		`DELETE FROM subscriptions WHERE user_id=$1 and segment_id=ANY($2)`,
		subs.IDUser, subs.IDSegment,
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
