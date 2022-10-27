package baserequest

type CreateRequestComment struct {
	Message string `json:"message"`
	PhotoId uint   `json:"photo_id"`
}
type UpdateRequestComment struct {
	Message string `json:"message"`
}
