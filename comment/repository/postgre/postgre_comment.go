package postgre

import (
	"final-project-go-clean-arch/domain"

	"gorm.io/gorm"
)

type commentRepository struct {
	DB *gorm.DB
}

func NewCommentRepository(db *gorm.DB) domain.CommentRepository {
	return &commentRepository{
		DB: db,
	}
}

// DeleteComment implements domain.CommentRepository
func (*commentRepository) DeleteComment(id string) error {
	panic("unimplemented")
}

// FindAll implements domain.CommentRepository
func (*commentRepository) FindAll() ([]*domain.Comment, error) {
	panic("unimplemented")
}

// FindById implements domain.CommentRepository
func (*commentRepository) FindById(id string) (*domain.Comment, error) {
	panic("unimplemented")
}

// Store implements domain.CommentRepository
func (*commentRepository) Store(comment *domain.Comment) (*domain.Comment, error) {
	panic("unimplemented")
}

// UpdateComment implements domain.CommentRepository
func (*commentRepository) UpdateComment(comment *domain.Comment) (*domain.Comment, error) {
	panic("unimplemented")
}
