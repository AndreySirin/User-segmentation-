package server

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"golang.org/x/sync/errgroup"

	"github.com/AndreySirin/avito-backend-assignment-2023/internal/entity"
	"github.com/AndreySirin/avito-backend-assignment-2023/internal/service"
)

// srv is service interface
type srv interface {
	CreateUser(ctx context.Context, request service.CreateUserRequest) (uuid.UUID, error)
	UpdateUser(ctx context.Context, userUpdate service.UpdateUserRequest) error
	DeleteUser(ctx context.Context, id uuid.UUID) error
	GetUser(ctx context.Context, id uuid.UUID) (*entity.User, error)
	ListUsers(ctx context.Context) ([]entity.User, error)
	//
	CreateSegment(ctx context.Context, request service.CreateSegmentRequest) (uuid.UUID, error)
	GetSegment(ctx context.Context, id uuid.UUID) (*entity.Segment, error)
	DeleteSegment(ctx context.Context, id uuid.UUID) error
	ListSegment(ctx context.Context) ([]entity.Segment, error)
	UpdateSegment(ctx context.Context, request service.UpdateSegmentRequest) (err error)
	//
	InsertUserInSegments(ctx context.Context, sub *service.CreateSubscription) error
	DeleteUserInSegments(ctx context.Context, sub *service.CreateSubscription) error
}

type Server struct {
	lg         *slog.Logger
	httpServer *http.Server

	service srv
}

func New(lg *slog.Logger, addr string, service *service.Service, module string) *Server {
	s := &Server{
		lg:      lg.With("module", module),
		service: service,
	}

	r := chi.NewRouter()
	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			// user
			r.Post("/users", s.handleCreateUser)
			r.Get("/users", s.handleListUsers)
			r.Get("/users/{id}", s.handleGetUser)
			r.Put("/users/{id}", s.handleUpdateUser)
			r.Delete("/users/{id}", s.handleDeleteUser)
			// segment
			r.Post("/segments", s.handleCreateSegment)
			r.Get("/segments", s.handleListSegments)
			r.Get("/segments/{id}", s.handleGetSegment)
			r.Put("/segments/{id}", s.handleUpdateSegment)
			r.Delete("/segments/{id}", s.handleDeleteSegment)
			// subscription
			r.Post("/subscription", s.CreateSubscription)
			r.Delete("/subscription/{id}", s.DeleteSubscription)
		})
	})

	s.httpServer = &http.Server{
		Addr:              addr,
		Handler:           r,
		ReadHeaderTimeout: time.Second,
	}

	return s
}

func (s *Server) Run(ctx context.Context, duration time.Duration) error {
	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		<-ctx.Done()

		gfCtx, cancel := context.WithTimeout(context.Background(), duration)
		defer cancel()

		s.lg.Info("graceful shutdown")
		return s.httpServer.Shutdown(
			gfCtx,
		)
	})

	eg.Go(func() error {
		s.lg.Info("listen and serve", "addr", s.httpServer.Addr)

		if err := s.httpServer.ListenAndServe(); err != nil &&
			!errors.Is(err, http.ErrServerClosed) {
			return fmt.Errorf("run: %v", err)
		}

		return nil
	})

	return eg.Wait()
}
