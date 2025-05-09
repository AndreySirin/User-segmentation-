package server

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"net/http"

	"github.com/AndreySirin/avito-backend-assignment-2023/internal/service"
)

func (s *Server) handleCreateUser(w http.ResponseWriter, r *http.Request) {

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
	w.Header().Set("Content-Type", "application/json")
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

func (s *Server) handleUpdateUser(w http.ResponseWriter, r *http.Request) {

	var req updateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	idUser := chi.URLParam(r, "id")
	id, err := uuid.Parse(idUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	req.Id = id
	if err = req.valid(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	userRequest, err := req.toService()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = s.service.UpdateUser(r.Context(), userRequest)
	if err != nil {
		if errors.Is(err, service.ErrNotValid) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
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

func (s *Server) handleDeleteUser(w http.ResponseWriter, r *http.Request) {
	idUser := chi.URLParam(r, "id")
	id, err := uuid.Parse(idUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if id == uuid.Nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	err = s.service.DeleteUser(r.Context(), id)
	if err != nil {
		if errors.Is(err, service.ErrNotValid) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
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
func (s *Server) handleGetUsers(w http.ResponseWriter, r *http.Request) {
	idUser := chi.URLParam(r, "id")
	id, err := uuid.Parse(idUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if id == uuid.Nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	user, err := s.service.GetUser(r.Context(), id)
	if err != nil {
		if errors.Is(err, service.ErrNotValid) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func (s *Server) handleListUser(w http.ResponseWriter, r *http.Request) {
	users, err := s.service.ListUsers(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(users); err != nil {
	}
}
