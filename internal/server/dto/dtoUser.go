package dto

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/AndreySirin/avito-backend-assignment-2023/internal/service"
	"github.com/AndreySirin/avito-backend-assignment-2023/internal/validator"
)

var ErrNotValidDate = errors.New("not valid date")

type CreateUserRequest struct {
	FullName    string `json:"full_name"     validate:"required"`
	Gender      string `json:"gender"        validate:"required,oneof=male female"`
	DateOfBirth string `json:"date_of_birth" validate:"required"`
}

type UpdateUserRequest struct {
	ID          uuid.UUID `json:"id"            validate:"required"`
	FullName    string    `json:"full_name"     validate:"required"`
	Gender      string    `json:"gender"        validate:"required,oneof=male female"`
	DateOfBirth string    `json:"date_of_birth" validate:"required"`
}

func (r *CreateUserRequest) Valid() error {
	return validator.Validator.Struct(r)
}

func (r *CreateUserRequest) ToService() (service.CreateUserRequest, error) {
	dateOfBirth, err := time.Parse(time.DateOnly, r.DateOfBirth)
	if err != nil {
		return service.CreateUserRequest{}, fmt.Errorf("%w: %v", ErrNotValidDate, err)
	}

	reqToService := service.CreateUserRequest{
		FullName:    r.FullName,
		Gender:      r.Gender,
		DateOfBirth: dateOfBirth,
	}

	return reqToService, nil
}

func (u *UpdateUserRequest) Valid() error { return validator.Validator.Struct(u) }

func (u *UpdateUserRequest) ToService() (service.UpdateUserRequest, error) {
	dateOfBirth, err := time.Parse(time.DateOnly, u.DateOfBirth)
	if err != nil {
		return service.UpdateUserRequest{}, fmt.Errorf("%w: %v", ErrNotValidDate, err)
	}
	reqToService := service.UpdateUserRequest{
		ID:          u.ID,
		FullName:    u.FullName,
		Gender:      u.Gender,
		DateOfBirth: dateOfBirth,
	}
	return reqToService, nil
}
