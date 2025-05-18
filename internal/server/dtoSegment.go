package server

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/AndreySirin/avito-backend-assignment-2023/internal/service"
	"github.com/AndreySirin/avito-backend-assignment-2023/internal/validator"
)

type CreateSegmentRequest struct {
	Title       string `json:"title"         validate:"required"`
	Description string `json:"description"   validate:"required"`
	AutoUserPrc uint8  `json:"auto_user_prc" validate:"required,gte=0,lte=100"`
}

func (c *CreateSegmentRequest) Validate() error { return validator.Validator.Struct(c) }

func (c *CreateSegmentRequest) ToService() service.CreateSegmentRequest {
	return service.CreateSegmentRequest{
		Title:       c.Title,
		Description: c.Description,
		AutoUserPrc: c.AutoUserPrc,
	}
}

type UpdateSegmentRequest struct {
	ID          uuid.UUID `json:"id"            validate:"required"`
	Title       string    `json:"title"         validate:"required"`
	Description string    `json:"description"   validate:"required"`
	AutoUserPrc uint8     `json:"auto_user_prc" validate:"required,gte=0,lte=100"`
}

func (u *UpdateSegmentRequest) Validate() error {
	if u.ID == uuid.Nil {
		return fmt.Errorf("id cannot be nil")
	}
	return validator.Validator.Struct(u)
}

func (u *UpdateSegmentRequest) ToService() service.UpdateSegmentRequest {
	return service.UpdateSegmentRequest{
		ID:          u.ID,
		Title:       u.Title,
		Description: u.Description,
		AutoUserPrc: u.AutoUserPrc,
	}
}
