package service

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/AndreySirin/avito-backend-assignment-2023/internal/entity"
	"github.com/AndreySirin/avito-backend-assignment-2023/internal/validator"
)

type CreateSegmentRequest struct {
	Title       string `validate:"required"`
	Description string `validate:"required"`
	AutoUserPrc uint8  `validate:"required,gte=0,lte=100"`
}

func (c *CreateSegmentRequest) Validate() error { return validator.Validator.Struct(c) }

type UpdateSegmentRequest struct {
	ID          uuid.UUID `validate:"required"`
	Title       string    `validate:"required"`
	Description string    `validate:"required"`
	AutoUserPrc uint8     `validate:"required,gte=0,lte=100"`
}

func (u *UpdateSegmentRequest) Validate() error { return validator.Validator.Struct(u) }

func (s *Service) CreateSegment(
	ctx context.Context,
	request CreateSegmentRequest,
) (uuid.UUID, error) {
	err := request.Validate()
	if err != nil {
		return uuid.Nil, err
	}

	segment := entity.Segment{
		ID:          uuid.New(),
		Title:       request.Title,
		Description: request.Description,
		AutoUserPrc: request.AutoUserPrc,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	id, err := s.repository.CreateSegment(ctx, segment)
	if err != nil {
		return uuid.Nil, err
	}
	return id, nil
}

func (s *Service) GetSegment(ctx context.Context, id uuid.UUID) (*entity.Segment, error) {
	segment, err := s.repository.GetSegment(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error servis %v", err)
	}
	return segment, nil
}

func (s *Service) ListSegment(ctx context.Context) ([]entity.Segment, error) {
	segments, err := s.repository.ListSegments(ctx)
	if err != nil {
		return nil, fmt.Errorf("error servis %v", err)
	}
	return segments, nil
}

func (s *Service) UpdateSegment(ctx context.Context, request UpdateSegmentRequest) (err error) {
	err = request.Validate()
	if err != nil {
		return fmt.Errorf("%w,%v", ErrNotValid, err)
	}
	segment := entity.Segment{
		ID:          request.ID,
		Title:       request.Title,
		Description: request.Description,
		AutoUserPrc: request.AutoUserPrc,
	}
	err = s.repository.UpDateSegment(ctx, segment)
	if err != nil {
		return fmt.Errorf("error servis %v", err)
	}
	return nil
}

func (s *Service) DeleteSegment(ctx context.Context, id uuid.UUID) error {
	err := s.repository.DeleteSegment(ctx, id)
	if err != nil {
		return fmt.Errorf("error DeleteSegment %v", err)
	}
	return nil
}
