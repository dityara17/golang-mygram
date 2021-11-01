package modeluser

import "time"

type Request struct {
	ID       uint   `json:"id,omitempty"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      int    `json:"age"`
}

type Response struct {
	ID        uint      `json:"id"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Age       int       `json:"age"`
}

type RequestLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResponseLogin struct {
	Token string `json:"token"`
}
