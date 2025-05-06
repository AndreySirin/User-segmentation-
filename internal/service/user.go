package service

import (
	"context"
	"github.com/AndreySirin/avito-backend-assignment-2023/internal/entity"
	"github.com/AndreySirin/avito-backend-assignment-2023/internal/logger"
)

type UserStorage interface {
	CreateUser(context.Context, entity.User) (int, error)
	UpdateUser(context.Context, entity.User) (err error)
	DeleteUser(context.Context, entity.User) error
}
type UserService struct {
	lg      logger.MyloggerInterface
	storage UserStorage
}

func NewUserService(lg *logger.MyLogger, storage UserStorage) *UserService {
	return &UserService{
		lg:      lg,
		storage: storage}
}

func (s *UserService) CreateUsers(ctx context.Context, user entity.User) (int, error) {
	return s.storage.CreateUser(ctx, user)
}
func (s *UserService) UpdateUsers(ctx context.Context, user entity.User) (err error) {
	return s.storage.UpdateUser(ctx, user)
}
func (s *UserService) DeleteUsers(ctx context.Context, user entity.User) error {
	return s.storage.DeleteUser(ctx, user)
}
