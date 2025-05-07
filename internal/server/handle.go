package server

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/google/uuid"

	"github.com/AndreySirin/avito-backend-assignment-2023/internal/service"
)

func (s *Server) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req createUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := req.valid(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userRequest, err := req.toService()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := s.service.CreateUser(r.Context(), userRequest)
	if err != nil {
		if errors.Is(err, service.ErrNotValid) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

	type response struct {
		ID uuid.UUID `json:"id"`
	}

	if err = json.NewEncoder(w).Encode(response{
		ID: id,
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
