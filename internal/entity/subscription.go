package entity

import "time"

type Subscription struct {
	ID        int       `json:"id"`
	IDUser    int       `json:"id_user"`
	IDSegment int       `json:"id_segment"`
	TTL       time.Time `json:"ttl"`
	AutoAdded bool      `json:"auto_added"`
	Created   time.Time `json:"created"`
	Updated   time.Time `json:"updated"`
	Deleted   time.Time `json:"deleted"`
}

type CreateSubscription struct {
	IdUser      int       `json:"id_user"`
	NameSegment []string  `json:"name_segment"`
	IdSegment   []int     `json:"id_segment"`
	TTL         time.Time `json:"ttl"`
	AutoAdded   bool      `json:"auto_added"`
	CreatedAt   time.Time `json:"created_at"`
	Updated     time.Time `json:"updated"`
	Deleted     time.Time `json:"deleted"`
}
