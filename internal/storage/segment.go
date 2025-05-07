package storage

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/AndreySirin/avito-backend-assignment-2023/internal/entity"
)

type SegmentStorage struct {
	lg *slog.Logger
	db *sql.DB
}

func NewSegment(db *Storage) *SegmentStorage {
	return &SegmentStorage{
		lg: db.lg,
		db: db.db,
	}
}

func (s *SegmentStorage) CreateSegment(ctx context.Context, segment entity.Segment) (int, error) {
	err := s.db.QueryRowContext(ctx,
		`INSERT INTO segments (title,description,auto_user_prc)
		VALUES ($1,$2,$3) RETURNING id_segment`, segment.Title, segment.Description, segment.AutoUserPrc).
		Scan(&segment.IDSegment)
	if err != nil {
		return 0, fmt.Errorf("error from CreateSegment %v", err)
	}
	return segment.IDSegment, nil
}

func (s *SegmentStorage) DeleteSegment(ctx context.Context, segment entity.Segment) error {
	res, err := s.db.ExecContext(
		ctx,
		"DELETE FROM segments WHERE id_segment=$1",
		segment.IDSegment,
	)
	if err != nil {
		return fmt.Errorf("error from DeleteSegment %w", err)
	}
	rows, err := res.RowsAffected()
	if err != nil || rows == 0 {
		return fmt.Errorf("failed to check affected rows: %w", err)
	}
	return nil
}

func (s *SegmentStorage) UpDateSegment(ctx context.Context, segment entity.Segment) (err error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("begin tx: %v", err)
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	var exists bool
	err = tx.QueryRowContext(ctx, "SELECT EXISTS(SELECT 1 FROM segments WHERE id_segment=$1)", segment.ID).
		Scan(&exists)
	if err != nil {
		return fmt.Errorf("check segment existence: %w", err)
	}
	if !exists {
		return fmt.Errorf("segment id %d does not exist", segment.ID)
	}

	_, err = tx.ExecContext(
		ctx,
		"UPDATE segments SET title=$1,description=$2,auto_user_prc=$3 WHERE id_segment=$4",
		segment.Title,
		segment.Description,
		segment.AutoUserPrc,
		segment.ID,
	)
	if err != nil {
		return fmt.Errorf("update segment: %w", err)
	}
	return nil
}
