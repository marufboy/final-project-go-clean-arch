package domain

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	Username  string    `json:"username" gorm:"NOT NULL;unique;type:varchar(255);" valid:"required"`
	Email     string    `json:"email" gorm:"NOT NULL;unique;type:varchar(255);" valid:"required,email"`
	Password  string    `json:"password,omitempty" gorm:"NOT NULL;type:text;" valid:"required,minstringlength(6)"`
	Age       uint      `json:"age" gorm:"NOT NULL;type:integer;" valid:"required,range(17|100)"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
