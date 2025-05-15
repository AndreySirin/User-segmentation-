package entity

import (
	"time"

	"github.com/google/uuid"
)

type Subscription struct {
	IDUser      uuid.UUID   `validate:"required"`
	IDSegment   []uuid.UUID `validate:"required"`
	NameSegment []string    `validate:"required,dive,required"`
	TTL         []time.Time
	AutoAdded   []bool
	Created     time.Time `validate:"required"`
	Updated     time.Time `validate:"required"`
	Deleted     time.Time
}
