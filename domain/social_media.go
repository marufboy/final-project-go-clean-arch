package domain

import (
	baserequest "final-project-go-clean-arch/common/base_request"
	"final-project-go-clean-arch/helpers"
	"time"
)

type SocialMedia struct {
	ID             uint      `json:"id" gorm:"primarykey"`
	Name           string    `json:"name" gorm:"NOT NULL;type:varchar(255);"`
	SocialMediaUrl string    `json:"social_media_url" gorm:"NOT NULL;type:text;"`
	UserID         uint      `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type SocialMediaUseCase interface {
	CreateSocialMediaUC(request baserequest.CreateRequestSocialMedia) helpers.Response
	UpdateSocialMediaUC(request baserequest.UpdateRequestComment) helpers.Response
	DeleteSocialMediaUC(id string) helpers.Response
	GetSocialMediaUC() helpers.Response
}
type SocialMediaRepository interface {
	Store(socialMedia *SocialMedia) (*SocialMedia, error)
	UpdateSocialMedia(socialMedia *SocialMedia) (*SocialMedia, error)
	DeleteSocialMedia(id string) error
	FindById(id string) (*SocialMedia, error)
	FindAll() ([]*SocialMedia, error)
}
