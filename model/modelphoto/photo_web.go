package modelphoto

import (
	"time"
)

type Request struct {
	Title    string `json:"title"`
	Caption  string `json:"caption,omitempty"`
	PhotoURL string `json:"photo_url"`
	UserID   uint   `json:"user_id"`
}

type Response struct {
	ID        uint      `json:"id,omitempty"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption,omitempty"`
	PhotoURL  string    `json:"photo_url"`
	UserID    uint      `json:"user_id,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type ResponseGet struct {
	ID       uint   `json:"id,omitempty"`
	Title    string `json:"title"`
	Caption  string `json:"caption,omitempty"`
	PhotoURL string `json:"photo_url"`
	User     struct {
		Username string `json:"username"`
		Email    string `json:"email"`
	} `json:"user"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type ResponseUpdate struct {
	ID        uint      `json:"id,omitempty"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption,omitempty"`
	PhotoURL  string    `json:"photo_url"`
	UserID    uint      `json:"user_id,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
