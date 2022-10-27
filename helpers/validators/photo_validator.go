package validators

import (
	baserequest "final-project-go-clean-arch/common/base_request"

	validation "github.com/asaskevich/govalidator"
)

func ValidateCreatePhoto(request baserequest.CreateRequestPhoto) error {
	_, errRequest := validation.ValidateStruct(&request)

	if errRequest != nil {
		return errRequest
	}
	return nil
}

func ValidateUpdatePhoto(request baserequest.UpdateRequestPhoto) error {
	_, errRequest := validation.ValidateStruct(&request)

	if errRequest != nil {
		return errRequest
	}
	return nil
}
