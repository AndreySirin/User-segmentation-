package server

import (
	"context"
	"encoding/json"
	"github.com/AndreySirin/avito-backend-assignment-2023/internal/entity"
	"github.com/AndreySirin/avito-backend-assignment-2023/internal/logger"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

type UserHandle interface {
	CreateUsers(context.Context, entity.User) (int, error)
	UpdateUsers(context.Context, entity.User) (err error)
	DeleteUsers(context.Context, entity.User) error
}

type SegmentHandle interface {
	CreateSegments(context.Context, entity.Segment) (int, error)
	DeleteSegments(context.Context, entity.Segment) error
	UpDateSegments(context.Context, entity.Segment) (err error)
}

type SubscriptionHandle interface {
	InsertUserInSegments(ctx context.Context, subs entity.CreateSubscription) (err error)
	DeleteUserInSegments(ctx context.Context, subs entity.CreateSubscription) (err error)
}

type HNDL struct {
	UserHandler         UserHandle
	SegmentHandler      SegmentHandle
	SubscriptionHandler SubscriptionHandle
	lg                  logger.MyloggerInterface
}

func NewHNDL(lg *logger.MyLogger, userHandler UserHandle, segmentHandler SegmentHandle, subscriptionHandler SubscriptionHandle) *HNDL {
	return &HNDL{
		lg:                  lg,
		UserHandler:         userHandler,
		SegmentHandler:      segmentHandler,
		SubscriptionHandler: subscriptionHandler,
	}
}

func (H *HNDL) UserAddSegment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	var subs entity.CreateSubscription
	err := json.NewDecoder(r.Body).Decode(&subs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = H.SubscriptionHandler.InsertUserInSegments(context.Background(), subs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode("successfully"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (H *HNDL) UserDeleteSegment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()
	var subs entity.CreateSubscription
	err := json.NewDecoder(r.Body).Decode(&subs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = H.SubscriptionHandler.DeleteUserInSegments(context.Background(), subs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode("successfully"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (H *HNDL) UserCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
	defer r.Body.Close()
	var user entity.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	id, err := H.UserHandler.CreateUsers(context.Background(), user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (H *HNDL) UserUpdate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	idUser := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var user entity.User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user.ID = id
	err = H.UserHandler.UpdateUsers(context.Background(), user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode("successfully"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func (H *HNDL) UserDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	idUser := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = H.UserHandler.DeleteUsers(context.Background(), entity.User{ID: id})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode("delete user with" + idUser); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (H *HNDL) SegmentCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()
	var segment entity.Segment
	err := json.NewDecoder(r.Body).Decode(&segment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id, err := H.SegmentHandler.CreateSegments(context.Background(), segment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func (H *HNDL) SegmentUpdate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	idSegment := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idSegment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var segment entity.Segment
	err = json.NewDecoder(r.Body).Decode(&segment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	segment.ID = id
	err = H.SegmentHandler.UpDateSegments(context.Background(), segment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode("successfully"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func (H *HNDL) SegmentDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	idSegment := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idSegment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = H.SegmentHandler.DeleteSegments(context.Background(), entity.Segment{IDSegment: id})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode("delete segment with" + idSegment); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
