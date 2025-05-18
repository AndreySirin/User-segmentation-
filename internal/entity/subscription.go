package entity

import (
	"time"

	"github.com/google/uuid"
)

type Subscription struct {
	UserID       uuid.UUID   `validate:"required"`
	SegmentID    []uuid.UUID `validate:"required"`
	TitleSegment []string    `validate:"required,dive,required"`
	TTL          []time.Time
	IsAutoAdded  bool
	CreatedAt    time.Time `validate:"required"`
	UpdatedAt    time.Time `validate:"required"`
	DeletedAt    *time.Time
}
