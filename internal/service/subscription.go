package service

import (
	"context"
	"github.com/AndreySirin/avito-backend-assignment-2023/internal/entity"
	"github.com/AndreySirin/avito-backend-assignment-2023/internal/logger"
)

type SubscriptionStorage interface {
	InsertUserInSegment(context.Context, entity.CreateSubscription) (err error)
	DeleteUserInSegment(context.Context, entity.CreateSubscription) (err error)
}

type SubscriptionService struct {
	lg      logger.MyloggerInterface
	storage SubscriptionStorage
}

func NewSubscriptionService(lg *logger.MyLogger, storage SubscriptionStorage) *SubscriptionService {
	return &SubscriptionService{
		lg:      lg,
		storage: storage}
}
func (s *SubscriptionService) InsertUserInSegments(ctx context.Context, subs entity.CreateSubscription) (err error) {
	return s.storage.InsertUserInSegment(ctx, subs)
}
func (s *SubscriptionService) DeleteUserInSegments(ctx context.Context, subs entity.CreateSubscription) (err error) {
	return s.storage.DeleteUserInSegment(ctx, subs)
}
