package validators

import (
	baserequest "final-project-go-clean-arch/common/base_request"

	validation "github.com/asaskevich/govalidator"
)

func ValidateCreateComment(request baserequest.CreateRequestComment) error {
	_, errRequest := validation.ValidateStruct(&request)

	if errRequest != nil {
		return errRequest
	}
	return nil
}

func ValidateUpdateComment(request baserequest.UpdateRequestComment) error {
	_, errRequest := validation.ValidateStruct(&request)

	if errRequest != nil {
		return errRequest
	}
	return nil
}
