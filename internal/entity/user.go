package entity

import (
	"time"

	"github.com/google/uuid"

	"github.com/AndreySirin/avito-backend-assignment-2023/internal/validator"
)

type User struct {
	ID uuid.UUID `validate:"required"`

	FullName    string    `validate:"required"`
	Gender      string    `validate:"required,oneof=male female"`
	DateOfBirth time.Time `validate:"required"`

	CreatedAt time.Time  `validate:"required"`
	UpdatedAt time.Time  `validate:"required"`
	DeletedAt *time.Time `validate:"omitempty"`
}

func (u *User) Validate() error {
	return validator.Validator.Struct(u)
}
