package baserequest

type CreateRequestSocialMedia struct {
	Name           string `json:"name" form:"name" validate:"required"`
	SocialMediaUrl string `json:"social_media_url" form:"social_media_url" validate:"required"`
}
type UpdateRequestSocialMedia struct {
	Name           string `json:"name" form:"name" validate:"required"`
	SocialMediaUrl string `json:"social_media_url" form:"social_media_url" validate:"required,url"`
}
