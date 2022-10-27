package baserequest

type CreateRequestPhoto struct {
	Title    string `json:"title" form:"title" validate:"required"`
	Caption  string `json:"caption" form:"caption"`
	PhotoUrl string `json:"photo_url" form:"photo_url" validate:"required"`
}
type UpdateRequestPhoto struct {
	Title    string `json:"title" form:"title" validate:"required"`
	Caption  string `json:"caption" form:"caption"`
	PhotoUrl string `json:"photo_url" form:"photo_url" validate:"required"`
}
