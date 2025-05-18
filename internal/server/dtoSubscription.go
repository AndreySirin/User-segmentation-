package server

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/AndreySirin/avito-backend-assignment-2023/internal/service"
	"github.com/AndreySirin/avito-backend-assignment-2023/internal/validator"
)

type SubscriptionRequest struct {
	UserID       uuid.UUID `json:"id_user"      validate:"required"`
	TitleSegment []string  `json:"name_segment" validate:"required,dive,required"`
	TTL          []string  `json:"ttl"`
	IsAutoAdded  bool      `json:"auto_added"`
}

func (s *SubscriptionRequest) Validate() error {
	if s.UserID == uuid.Nil {
		return errors.New("id_user is required")
	}
	return validator.Validator.Struct(s)
}

func (s *SubscriptionRequest) ToService() (*service.CreateSubscription, error) {
	var ttlSlice []time.Time
	for i := 0; i < len(s.TTL); i++ {
		date, err := time.Parse("2006-01-02 15:04:05", s.TTL[i])
		if err != nil {
			return nil, fmt.Errorf("invalid TTL at index %d: %w", i, err)
		}
		ttlSlice = append(ttlSlice, date)
	}
	subs := &service.CreateSubscription{
		UserID:       s.UserID,
		TitleSegment: s.TitleSegment,
		TTL:          ttlSlice,
		IsAutoAdded:  s.IsAutoAdded,
	}
	return subs, nil
}
