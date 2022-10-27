package validators

import (
	baserequest "final-project-go-clean-arch/common/base_request"

	validation "github.com/asaskevich/govalidator"
)

func ValidateRegister(request baserequest.RegisterRequest) error {
	_, errRequest := validation.ValidateStruct(&request)

	if errRequest != nil {
		return errRequest
	}
	return nil
}

func ValidateLogin(request baserequest.LoginRequest) error {
	_, errRequest := validation.ValidateStruct(&request)

	if errRequest != nil {
		return errRequest
	}
	return nil
}

func ValidateUpdateUser(request baserequest.UpdateRequestUser) error {
	_, errRequest := validation.ValidateStruct(&request)

	if errRequest != nil {
		return errRequest
	}
	return nil
}
