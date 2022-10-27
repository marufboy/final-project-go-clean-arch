package baseresponse

import "time"

type UserResponse struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Age       uint      `json:"age"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
