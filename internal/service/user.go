package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/AndreySirin/avito-backend-assignment-2023/internal/entity"
	"github.com/AndreySirin/avito-backend-assignment-2023/internal/storage"
	"github.com/AndreySirin/avito-backend-assignment-2023/internal/validator"
)

var ErrNotValid = errors.New("bad request")

type CreateUserRequest struct {
	FullName    string    `validate:"required"`
	Gender      string    `validate:"required,oneof=male female"`
	DateOfBirth time.Time `validate:"required"`
}

func (r *CreateUserRequest) Validate() error {
	return validator.Validator.Struct(r)
}

func (s *Service) CreateUser(
	ctx context.Context,
	createUserRequest CreateUserRequest,
) (uuid.UUID, error) {
	if err := createUserRequest.Validate(); err != nil {
		return uuid.Nil, fmt.Errorf("%w: %v", ErrNotValid, err)
	}

	t := time.Now()

	user := entity.User{
		ID:          uuid.New(),
		FullName:    createUserRequest.FullName,
		Gender:      createUserRequest.Gender,
		DateOfBirth: createUserRequest.DateOfBirth,
		CreatedAt:   t,
		UpdatedAt:   t,
		DeletedAt:   nil,
	}

	id, err := s.repository.CreateUser(ctx, user)
	if err != nil {
		if errors.Is(err, storage.ErrNotValid) {
			return uuid.Nil, fmt.Errorf("%w: %v", ErrNotValid, err)
		}

		return uuid.Nil, fmt.Errorf("CreateUser: %w", err)
	}

	return id, nil
}

// TODO
// - GetUser
// - ListUsers

func (s *Service) UpdateUser(ctx context.Context, user entity.User) error {
	// FIXME
	return nil
}

func (s *Service) DeleteUser(ctx context.Context, user entity.User) error {
	// FIXME
	return nil
}
