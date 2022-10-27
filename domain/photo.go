package domain

import "time"

type Photo struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	Title     string    `json:"title" gorm:"NOT NULL;type:varchar(255);" valid:"required"`
	Caption   string    `json:"caption" gorm:"type:varchar(255);"`
	PhotoUrl  string    `json:"photo_url" gorm:"NOT NULL;type:text;" valid:"required"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
