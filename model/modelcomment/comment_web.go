package modelcomment

import "time"

type Request struct {
	Message string `json:"message"`
	PhotoID uint   `json:"photo_id,omitempty"`
	UserID  uint   `json:"user_id,omitempty"`
}

type RequestUpdate struct {
	Message string `json:"message"`
	UserID  uint   `json:"user_id,omitempty"`
}

type ResponseInsert struct {
	ID        uint      `json:"id"`
	Message   string    `json:"message"`
	PhotoID   uint      `json:"photo_id"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type ResponseUpdate struct {
	ID        uint      `json:"id"`
	Message   string    `json:"message"`
	PhotoID   uint      `json:"photo_id"`
	UserID    uint      `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Response struct {
	ID        uint      `json:"id"`
	Message   string    `json:"message"`
	PhotoID   uint      `json:"photo_id"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	User      struct {
		ID       uint   `json:"id"`
		Email    string `json:"email"`
		Username string `json:"username"`
	} `json:"user"`
	Photo struct {
		ID       uint   `json:"id"`
		Title    string `json:"title"`
		Caption  string `json:"caption,omitempty"`
		PhotoURL string `json:"photo_url"`
		UserID   uint   `json:"user_id"`
	} `json:"photo"`
}
