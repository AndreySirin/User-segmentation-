package dto

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/AndreySirin/avito-backend-assignment-2023/internal/service"
	"github.com/AndreySirin/avito-backend-assignment-2023/internal/validator"
)

type SubscriptionRequest struct {
	IDUser      uuid.UUID `json:"id_user"      validate:"required"`
	NameSegment []string  `json:"name_segment" validate:"required,dive,required"`
	TTL         []string  `json:"ttl"`
	AutoAdded   []bool    `json:"auto_added"`
}

func (s *SubscriptionRequest) Validate() error {
	if s.IDUser == uuid.Nil {
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
		IDUser:      s.IDUser,
		NameSegment: s.NameSegment,
		TTL:         ttlSlice,
		AutoAdded:   s.AutoAdded,
	}
	return subs, nil
}
