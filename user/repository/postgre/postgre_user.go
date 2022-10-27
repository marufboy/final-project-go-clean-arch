package postgre

import (
	"final-project-go-clean-arch/domain"
	"final-project-go-clean-arch/helpers"

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
func (r *userRepository) FindByEmail(email string) (*domain.User, error) {
	var user *domain.User
	err := r.DB.Where("email=?", email).Take(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

// FindByUsername implements domain.UserRepository
func (r *userRepository) FindByUsername(username string) (*domain.User, error) {
	var user *domain.User
	err := r.DB.Where("username=?", username).Take(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Store implements domain.UserRepository
func (r *userRepository) Store(user *domain.User) (*domain.User, error) {
	//TODO: replace down code in usecase
	user.Password = helpers.EncryptPassword(user.Password)
	err := r.DB.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

// UpdateUser implements domain.UserRepository
func (r *userRepository) UpdateUser(user *domain.User) (*domain.User, error) {
	err := r.DB.Model(&user).Updates(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

// DeleteUser implements domain.UserRepository
func (r *userRepository) DeleteUser(id string) error {
	var oldUser *domain.User
	err := r.DB.First(&oldUser, "id=?", id).Error
	if err != nil {
		return nil
	}

	return err
}
