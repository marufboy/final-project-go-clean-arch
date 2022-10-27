package postgre

import (
	"final-project-go-clean-arch/domain"

	"gorm.io/gorm"
)

type photoRepository struct {
	DB *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) domain.PhotoRepository {
	return &photoRepository{
		DB: db,
	}
}

// DeletePhoto implements domain.PhotoRepository
func (*photoRepository) DeletePhoto(id string) error {
	panic("unimplemented")
}

// FindAll implements domain.PhotoRepository
func (*photoRepository) FindAll() ([]*domain.Photo, error) {
	panic("unimplemented")
}

// FindById implements domain.PhotoRepository
func (*photoRepository) FindById(id string) (*domain.Photo, error) {
	panic("unimplemented")
}

// Store implements domain.PhotoRepository
func (*photoRepository) Store(photo *domain.Photo) (*domain.Photo, error) {
	panic("unimplemented")
}

// UpdatePhoto implements domain.PhotoRepository
func (*photoRepository) UpdatePhoto(photo *domain.Photo) (*domain.Photo, error) {
	panic("unimplemented")
}
