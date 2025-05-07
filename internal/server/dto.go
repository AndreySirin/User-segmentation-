package server

import (
	"errors"
	"fmt"
	"time"

	"github.com/AndreySirin/avito-backend-assignment-2023/internal/service"
	"github.com/AndreySirin/avito-backend-assignment-2023/internal/validator"
)

type createUserRequest struct {
	FullName    string `json:"fullName"    validate:"required"`
	Gender      string `json:"gender"      validate:"required,oneof=male female"`
	DateOfBirth string `json:"dateOfBirth" validate:"required"`
}

func (r *createUserRequest) valid() error {
	return validator.Validator.Struct(r)
}

var ErrNotValidDate = errors.New("not valid date")

func (r *createUserRequest) toService() (service.CreateUserRequest, error) {
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
