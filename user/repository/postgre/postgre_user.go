package postgre

import (
	"final-project-go-clean-arch/domain"

	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{
		DB: db,
	}
}

// FindByEmail implements domain.UserRepository
func (*userRepository) FindByEmail(email string) (*domain.User, error) {
	panic("unimplemented")
}

// FindByUsername implements domain.UserRepository
func (*userRepository) FindByUsername(username string) (*domain.User, error) {
	panic("unimplemented")
}

// Store implements domain.UserRepository
func (*userRepository) Store(user domain.User) (*domain.User, error) {
	panic("unimplemented")
}

// UpdateUser implements domain.UserRepository
func (*userRepository) UpdateUser(user domain.User) (*domain.User, error) {
	panic("unimplemented")
}
