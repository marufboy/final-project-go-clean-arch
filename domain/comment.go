package domain

import (
	baserequest "final-project-go-clean-arch/common/base_request"
	"final-project-go-clean-arch/helpers"
	"time"
)

type Comment struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	Message   string    `json:"message" gorm:"NOT NULL;type:text;" valid:"required"`
	UserID    uint      `json:"user_id"`
	PhotoID   uint      `json:"photo_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CommentUseCase interface {
	CreateCommentUC(request baserequest.CreateRequestComment) helpers.Response
	UpdateCommentUC(request baserequest.UpdateRequestComment) helpers.Response
	DeleteCommentUC(id string) helpers.Response
	GetCommentUC() helpers.Response
}
type CommentRepository interface {
	Store(comment *Comment) (*Comment, error)
	UpdateComment(comment *Comment) (*Comment, error)
	DeleteComment(id string) error
	FindById(id string) (*Comment, error)
	FindAll() ([]*Comment, error)
}
