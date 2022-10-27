package usecase

import (
	baserequest "final-project-go-clean-arch/common/base_request"
	"final-project-go-clean-arch/domain"
	"final-project-go-clean-arch/helpers"
)

type photoUseCase struct {
	photoRepository domain.PhotoRepository
}

func NewPhotoUseCase(repository domain.PhotoRepository) domain.PhotoUseCase {
	return &photoUseCase{
		photoRepository: repository,
	}
}

// CreatePhotoUC implements domain.PhotoUseCase
func (*photoUseCase) CreatePhotoUC(request baserequest.CreateRequestPhoto) helpers.Response {
	panic("unimplemented")
}

// DeletePhotoUC implements domain.PhotoUseCase
func (*photoUseCase) DeletePhotoUC(id string) helpers.Response {
	panic("unimplemented")
}

// GetPhotoUC implements domain.PhotoUseCase
func (*photoUseCase) GetPhotoUC() helpers.Response {
	panic("unimplemented")
}

// UpdatePhotoUC implements domain.PhotoUseCase
func (*photoUseCase) UpdatePhotoUC(request baserequest.UpdateRequestPhoto) helpers.Response {
	panic("unimplemented")
}
