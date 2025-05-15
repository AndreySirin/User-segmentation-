package service

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/AndreySirin/avito-backend-assignment-2023/internal/entity"
)

type CreateSubscription struct {
	IDUser      uuid.UUID `validate:"required"`
	NameSegment []string  `validate:"required,dive,required"`
	IDSegment   []uuid.UUID
	TTL         []time.Time
	AutoAdded   []bool
}

func (s *Service) InsertUserInSegments(ctx context.Context, req *CreateSubscription) error {
	tx, err := s.repository.TX(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	subscription := &entity.Subscription{
		IDUser:      req.IDUser,
		NameSegment: req.NameSegment,
		TTL:         req.TTL,
		AutoAdded:   req.AutoAdded,
	}

	err = s.repository.CheckExistUser(ctx, tx, subscription)
	if err != nil {
		return err
	}
	subscription.IDSegment, err = s.repository.GetIDForSegment(ctx, tx, subscription)
	if err != nil {
		return err
	}

	err = s.repository.InsertSubscription(ctx, tx, subscription)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) DeleteUserInSegments(ctx context.Context, req *CreateSubscription) error {
	tx, err := s.repository.TX(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	subscription := &entity.Subscription{
		IDUser:      req.IDUser,
		NameSegment: req.NameSegment,
		TTL:         req.TTL,
		AutoAdded:   req.AutoAdded,
	}

	subscription.IDSegment, err = s.repository.GetIDForSegment(ctx, tx, subscription)
	if err != nil {
		return err
	}
	err = s.repository.DeleteSubscription(ctx, tx, subscription)
	if err != nil {
		return err
	}
	return nil
}
