package service

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/AndreySirin/avito-backend-assignment-2023/internal/entity"
)

type CreateSubscription struct {
	UserID       uuid.UUID `validate:"required"`
	TitleSegment []string  `validate:"required,dive,required"`
	SegmentID    []uuid.UUID
	TTL          []time.Time
	IsAutoAdded  bool
}

func (s *Service) InsertUserInSegments(ctx context.Context, req *CreateSubscription) (err error) {
	tx, err := s.repository.TX(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			errRollback := tx.Rollback()
			if errRollback != nil {
				s.lg.Error("error to rollback transaction", "err", errRollback)
			}
		} else {
			errCommit := tx.Commit()
			if errCommit != nil {
				s.lg.Error("error to commit transaction", "err", errCommit)
				err = errCommit
			}
		}
	}()

	subscription := &entity.Subscription{
		UserID:       req.UserID,
		TitleSegment: req.TitleSegment,
		TTL:          req.TTL,
		IsAutoAdded:  req.IsAutoAdded,
	}

	err = s.repository.CheckExistUser(ctx, tx, subscription)
	if err != nil {
		return err
	}
	subscription.SegmentID, err = s.repository.GetIDForSegment(ctx, tx, subscription)
	if err != nil {
		return err
	}

	err = s.repository.InsertSubscription(ctx, tx, subscription)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) DeleteUserInSegments(ctx context.Context, req *CreateSubscription) (err error) {
	tx, err := s.repository.TX(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			errRollback := tx.Rollback()
			if errRollback != nil {
				s.lg.Error("error to rollback transaction", "err", errRollback)
			}
		} else {
			errCommit := tx.Commit()
			if errCommit != nil {
				s.lg.Error("error to commit transaction", "err", errCommit)
				err = errCommit
			}
		}
	}()

	subscription := &entity.Subscription{
		UserID:       req.UserID,
		TitleSegment: req.TitleSegment,
		TTL:          req.TTL,
		IsAutoAdded:  req.IsAutoAdded,
	}

	subscription.SegmentID, err = s.repository.GetIDForSegment(ctx, tx, subscription)
	if err != nil {
		return err
	}
	err = s.repository.DeleteSubscription(ctx, tx, subscription)
	if err != nil {
		return err
	}
	return nil
}
