package entity

import "time"

type Segment struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	AutoUserPrc uint8     `json:"auto_user_prc"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
	IDSegment   int       `json:"id_segment"`
}
