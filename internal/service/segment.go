package service

import (
	"context"

	"github.com/google/uuid"

	"github.com/AndreySirin/avito-backend-assignment-2023/internal/entity"
)

func (s *Service) CreateSegment(ctx context.Context, segment entity.Segment) (uuid.UUID, error) {
	// FIXME
	return uuid.Nil, nil
}

// TODO
// - GetSegment
// - ListSegments

func (s *Service) UpdateSegment(ctx context.Context, segment entity.Segment) (err error) {
	// FIXME
	return nil
}

func (s *Service) DeleteSegment(ctx context.Context, segment entity.Segment) error {
	// FIXME
	return nil
}
