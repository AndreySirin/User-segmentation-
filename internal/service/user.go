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

type UpdateUserRequest struct {
	ID          uuid.UUID `validate:"required"`
	FullName    string    `validate:"required"`
	Gender      string    `validate:"required,oneof=male female"`
	DateOfBirth time.Time `validate:"required"`
}

func (r *CreateUserRequest) Validate() error {
	return validator.Validator.Struct(r)
}

func (u *UpdateUserRequest) Validate() error {
	return validator.Validator.Struct(u)
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

func (s *Service) UpdateUser(ctx context.Context, userUpdate UpdateUserRequest) error {
	if err := userUpdate.Validate(); err != nil {
		return fmt.Errorf("%w: %v", ErrNotValid, err)
	}
	t := time.Now()
	user := entity.User{
		ID:          userUpdate.ID,
		FullName:    userUpdate.FullName,
		Gender:      userUpdate.Gender,
		DateOfBirth: userUpdate.DateOfBirth,
		UpdatedAt:   t,
	}
	err := s.repository.UpdateUser(ctx, user)
	if err != nil {
		return fmt.Errorf("calling the method of db s.repository.UpdateUser(ctx, user) : %w", err)
	}
	return nil
}

func (s *Service) GetUser(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	return s.repository.GetUser(ctx, id)
}

func (s *Service) ListUsers(ctx context.Context) ([]entity.User, error) {
	return s.repository.ListUsers(ctx)
}

func (s *Service) DeleteUser(ctx context.Context, id uuid.UUID) error {
	return s.repository.DeleteUser(ctx, id)
}
