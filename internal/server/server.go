package server

import (
	"fmt"
	"github.com/AndreySirin/avito-backend-assignment-2023/internal/logger"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	lg         *logger.MyLogger
	httpServer *http.Server
}

func NewServer(lg *logger.MyLogger, adr string, hndl *HNDL) *Server {
	s := &Server{
		lg: lg,
	}
	r := chi.NewRouter()
	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Post("/UserCreate", hndl.UserCreate)
			r.Put("/UserUpdate/{id}", hndl.UserUpdate)
			r.Delete("/UserDelete/{id}", hndl.UserDelete)

			r.Post("/SegmentCreate", hndl.SegmentCreate)
			r.Put("/SegmentUpdate/{id}", hndl.SegmentUpdate)
			r.Delete("/SegmentDelete/{id}", hndl.SegmentDelete)

			r.Post("/UserAddSegment", hndl.UserAddSegment)
			r.Delete("/DeleteUserSegment/{id}", hndl.UserDeleteSegment)
		})
	})

	s.httpServer = &http.Server{
		Addr:    adr,
		Handler: r,
	}
	return s
}

func (s *Server) Start() error {
	s.lg.Info(fmt.Sprintf("Starting server on %s", s.httpServer.Addr))
	return s.httpServer.ListenAndServe()
}
