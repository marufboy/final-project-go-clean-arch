package usecase

import (
	baserequest "final-project-go-clean-arch/common/base_request"
	baseresponse "final-project-go-clean-arch/common/base_response"
	"final-project-go-clean-arch/domain"
	"final-project-go-clean-arch/helpers"
	"time"

	"github.com/mashingan/smapping"
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
func (u *userUseCase) DeleteUserUC(id string) (res helpers.Response) {
	err := u.userRepository.DeleteUser(id)
	if err != nil {
		return helpers.BuildErrorResponse(err.Error(), helpers.EmptyObj{})
	}
	res = helpers.BuildResponse("Your account has been successfully deleted", helpers.EmptyObj{})
	return res
}

// UpdateUserUC implements domain.UserUseCase
func (u *userUseCase) UpdateUserUC(request baserequest.UpdateRequestUser) (res helpers.Response) {
	var user *domain.User
	err := smapping.FillStruct(&user, smapping.MapFields(&request))
	if err != nil {
		return helpers.BuildErrorResponse(err.Error(), helpers.EmptyObj{})
	}

	user.UpdatedAt = time.Now()
	user, err = u.userRepository.UpdateUser(user)
	if err != nil {
		return helpers.BuildErrorResponse(err.Error(), helpers.EmptyObj{})
	}

	data := baseresponse.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Age:       user.Age,
		UpdatedAt: user.UpdatedAt,
	}

	res = helpers.BuildResponse("Updated Sucessfully", data)
	return res
}

// UserLoginUC implements domain.UserUseCase
func (u *userUseCase) UserLoginUC(request baserequest.LoginRequest) (res helpers.Response) {
	var user *domain.User
	err := smapping.FillStruct(&user, smapping.MapFields(&request))
	if err != nil {
		return helpers.BuildErrorResponse(err.Error(), helpers.EmptyObj{})
	}

	//handling password user request
	password := user.Password

	user, err = u.userRepository.FindByEmail(user.Email)
	if err != nil {
		return helpers.BuildErrorResponse(err.Error(), helpers.EmptyObj{})
	}

	isValid := helpers.ComparePass(user.Password, []byte(password))
	if !isValid {
		res = helpers.BuildErrorResponse("Password invalid", helpers.EmptyObj{})
		return res
	}

	token := helpers.GenerateToken(user.ID)
	res = helpers.BuildResponse("Login succesfull", baseresponse.LoginResponse{Token: token})
	return res
}

// UserRegisterUC implements domain.UserUseCase
func (u *userUseCase) UserRegisterUC(request baserequest.RegisterRequest) (res helpers.Response) {
	var user *domain.User
	err := smapping.FillStruct(&user, smapping.MapFields(&request))
	if err != nil {
		return helpers.BuildErrorResponse(err.Error(), helpers.EmptyObj{})
	}

	user.Password = helpers.EncryptPassword(user.Password)
	user, err = u.userRepository.Store(user)
	if err != nil {
		return helpers.BuildErrorResponse(err.Error(), helpers.EmptyObj{})
	}

	data := baseresponse.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Age:      user.Age,
	}

	res = helpers.BuildResponse("Created Sucessfully", data)
	return res
}
