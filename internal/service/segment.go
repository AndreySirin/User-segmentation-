package service

import (
	"context"
	"github.com/AndreySirin/avito-backend-assignment-2023/internal/entity"
	"github.com/AndreySirin/avito-backend-assignment-2023/internal/logger"
)

type SegmentStorage interface {
	CreateSegment(context.Context, entity.Segment) (int, error)
	DeleteSegment(context.Context, entity.Segment) error
	UpDateSegment(context.Context, entity.Segment) (err error)
}
type SegmentServis struct {
	lg      logger.MyloggerInterface
	storage SegmentStorage
}

func NewSegment(lg *logger.MyLogger, storage SegmentStorage) *SegmentServis {
	return &SegmentServis{
		lg:      lg,
		storage: storage,
	}
}
func (s *SegmentServis) CreateSegments(ctx context.Context, segment entity.Segment) (int, error) {
	return s.storage.CreateSegment(ctx, segment)
}
func (s *SegmentServis) DeleteSegments(ctx context.Context, segment entity.Segment) error {
	return s.storage.DeleteSegment(ctx, segment)
}
func (s *SegmentServis) UpDateSegments(ctx context.Context, segment entity.Segment) (err error) {
	return s.storage.UpDateSegment(ctx, segment)
}
