package domain

import (
	baserequest "final-project-go-clean-arch/common/base_request"
	"final-project-go-clean-arch/helpers"
	"time"
)

type Photo struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	Title     string    `json:"title" gorm:"NOT NULL;type:varchar(255);" valid:"required"`
	Caption   string    `json:"caption" gorm:"type:varchar(255);"`
	PhotoUrl  string    `json:"photo_url" gorm:"NOT NULL;type:text;" valid:"required"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PhotoUseCase interface {
	CreatePhotoUC(request baserequest.CreateRequestPhoto) helpers.Response
	UpdatePhotoUC(request baserequest.UpdateRequestPhoto) helpers.Response
	DeletePhotoUC(id string) helpers.Response
	GetPhotoUC() helpers.Response
}
type PhotoRepository interface {
	Store(photo *Photo) (*Photo, error)
	UpdatePhoto(photo *Photo) (*Photo, error)
	DeletePhoto(id string) error
	FindById(id string) (*Photo, error)
	FindAll() ([]*Photo, error)
}
