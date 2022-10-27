package domain

import "time"

type Comment struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	Message   string    `json:"message" gorm:"NOT NULL;type:text;" valid:"required"`
	UserID    uint      `json:"user_id"`
	PhotoID   uint      `json:"photo_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
