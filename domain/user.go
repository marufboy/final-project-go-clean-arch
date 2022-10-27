package domain

import (
	baserequest "final-project-go-clean-arch/common/base_request"
	"final-project-go-clean-arch/helpers"
	"time"
)

type User struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	Username  string    `json:"username" gorm:"NOT NULL;unique;type:varchar(255);" valid:"required"`
	Email     string    `json:"email" gorm:"NOT NULL;unique;type:varchar(255);" valid:"required,email"`
	Password  string    `json:"password,omitempty" gorm:"NOT NULL;type:text;" valid:"required,minstringlength(6)"`
	Age       uint      `json:"age" gorm:"NOT NULL;type:integer;" valid:"required,range(8|100)"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserUseCase interface {
	UserRegisterUC(request baserequest.RegisterRequest) helpers.Response
	UserLoginUC(request baserequest.LoginRequest) helpers.Response
	UpdateUserUC(request baserequest.UpdateRequestUser) helpers.Response
	DeleteUserUC(id string) helpers.Response
}

type UserRepository interface {
	Store(user *User) (*User, error)
	UpdateUser(user *User) (*User, error)
	DeleteUser(id string) error
	FindByEmail(email string) (*User, error)
	FindByUsername(username string) (*User, error)
}
