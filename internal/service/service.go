package service

import (
	"context"
	"log/slog"

	"github.com/google/uuid"

	"github.com/AndreySirin/avito-backend-assignment-2023/internal/entity"
)

const module = "service"

type repository interface {
	CreateUser(ctx context.Context, user entity.User) (uuid.UUID, error)
	// TODO будут добавляться остальные методы из storage
	// ...
}

type Service struct {
	lg         *slog.Logger
	repository repository
}

func New(lg *slog.Logger, repository repository) *Service {
	return &Service{
		lg:         lg.With("module", module),
		repository: repository,
	}
}
