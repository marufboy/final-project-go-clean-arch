package usecase

import (
	baserequest "final-project-go-clean-arch/common/base_request"
	"final-project-go-clean-arch/domain"
	"final-project-go-clean-arch/helpers"
)

type commentUseCase struct {
	commentRepository domain.CommentRepository
}

func NewCommentUseCase(repository domain.CommentRepository) domain.CommentUseCase {
	return &commentUseCase{
		commentRepository: repository,
	}
}

// CreateCommentUC implements domain.CommentUseCase
func (*commentUseCase) CreateCommentUC(request baserequest.CreateRequestComment) helpers.Response {
	panic("unimplemented")
}

// DeleteCommentUC implements domain.CommentUseCase
func (*commentUseCase) DeleteCommentUC(id string) helpers.Response {
	panic("unimplemented")
}

// GetCommentUC implements domain.CommentUseCase
func (*commentUseCase) GetCommentUC() helpers.Response {
	panic("unimplemented")
}

// UpdateCommentUC implements domain.CommentUseCase
func (*commentUseCase) UpdateCommentUC(request baserequest.UpdateRequestComment) helpers.Response {
	panic("unimplemented")
}
