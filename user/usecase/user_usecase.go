package usecase

import (
	baserequest "final-project-go-clean-arch/common/base_request"
	"final-project-go-clean-arch/domain"
)

type userUseCase struct {
	userRepository domain.UserRepository
}

func NewUserUseCase(repository domain.UserRepository) domain.UserUseCase {
	return &userUseCase{
		userRepository: repository,
	}
}

// DeleteUserUC implements domain.UserUseCase
func (*userUseCase) DeleteUserUC(id string) (string, error) {
	panic("unimplemented")
}

// UpdateUserUC implements domain.UserUseCase
func (*userUseCase) UpdateUserUC(request baserequest.UpdateRequestUser) (*domain.User, error) {
	panic("unimplemented")
}

// UserLoginUC implements domain.UserUseCase
func (*userUseCase) UserLoginUC(request baserequest.LoginRequest) (string, error) {
	panic("unimplemented")
}

// UserRegisterUC implements domain.UserUseCase
func (*userUseCase) UserRegisterUC(request baserequest.RegisterRequest) (*domain.User, error) {
	panic("unimplemented")
}
