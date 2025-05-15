package entity

import (
	"time"

	"github.com/google/uuid"

	"github.com/AndreySirin/avito-backend-assignment-2023/internal/validator"
)

type Segment struct {
	ID          uuid.UUID `validate:"required"`
	Title       string    `validate:"required"`
	Description string    `validate:"required"`
	AutoUserPrc uint8     `validate:"required,gte=0,lte=100"`
	CreatedAt   time.Time `validate:"required"`
	UpdatedAt   time.Time `validate:"required"`
	DeletedAt   time.Time
}

func (s *Segment) Validate() error { return validator.Validator.Struct(s) }
