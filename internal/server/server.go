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

	"github.com/AndreySirin/avito-backend-assignment-2023/internal/service"
)

const (
	module = "server"

	// можно вынести в конфиг
	shutdownTimeout = 3 * time.Second
)

// srv is service interface
type srv interface {
	CreateUser(ctx context.Context, request service.CreateUserRequest) (uuid.UUID, error)
}

type Server struct {
	lg         *slog.Logger
	httpServer *http.Server

	service srv
}

func New(lg *slog.Logger, addr string, service srv) *Server {
	s := &Server{
		lg:      lg.With("module", module),
		service: service,
	}

	r := chi.NewRouter()
	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Post("/users", s.handleCreateUser)
			// TODO
			// r.Get("/users", s.handleListUsers)
			// r.Get("/users/{id}", s.handleGetUser)
			// r.Put("/users/{id}", s.handleUpdateUser)
			// r.Delete("/users/{id}", s.handleDeleteUser)
		})
	})

	s.httpServer = &http.Server{
		Addr:              addr,
		Handler:           r,
		ReadHeaderTimeout: time.Second,
	}

	return s
}

func (s *Server) Run(ctx context.Context) error {
	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		<-ctx.Done()

		gfCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
		defer cancel()

		s.lg.Info("graceful shutdown")
		return s.httpServer.Shutdown(
			gfCtx,
		) //nolint:contextcheck // graceful shutdown with new context
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
