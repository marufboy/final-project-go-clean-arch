package validators

import (
	baserequest "final-project-go-clean-arch/common/base_request"

	validation "github.com/asaskevich/govalidator"
)

func ValidateCreateSocialMedia(request baserequest.CreateRequestSocialMedia) error {
	_, errRequest := validation.ValidateStruct(&request)

	if errRequest != nil {
		return errRequest
	}
	return nil
}

func ValidateUpdateSocialMedia(request baserequest.UpdateRequestSocialMedia) error {
	_, errRequest := validation.ValidateStruct(&request)

	if errRequest != nil {
		return errRequest
	}
	return nil
}
