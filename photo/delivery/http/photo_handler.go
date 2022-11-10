package http

import (
	baserequest "final-project-go-clean-arch/common/base_request"
	"final-project-go-clean-arch/domain"
	"final-project-go-clean-arch/helpers"
	"final-project-go-clean-arch/helpers/validators"
	"final-project-go-clean-arch/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PhotoHandler struct {
	photoUseCase domain.PhotoUseCase
}

func NewPhotoHandler(r *gin.Engine, photoUC domain.PhotoUseCase, db *gorm.DB) {
	handler := &PhotoHandler{
		photoUseCase: photoUC,
	}

	router := r.Group("/photos", middlewares.AuthorizeJWT())
	{
		router.POST("/", handler.CreatePhoto)
		router.GET("/", handler.GetPhoto)
		router.PUT("/:photoId", middlewares.AuthorizationPhoto(db), handler.UpdatePhoto)
		router.DELETE("/:photoId", middlewares.AuthorizationPhoto(db), handler.DeletePhoto)
	}
}

func (h PhotoHandler) CreatePhoto(ctx *gin.Context) {
	var request baserequest.CreateRequestPhoto
	err := ctx.ShouldBind(&request)

	if err != nil {
		response := helpers.BuildErrorResponse(err.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	err = validators.ValidateCreatePhoto(request)
	if err != nil {
		response := helpers.BuildErrorResponse(err.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	res := h.photoUseCase.CreatePhotoUC(request)
	if res.Error {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
func (h PhotoHandler) GetPhoto(ctx *gin.Context) {
	res := h.photoUseCase.GetPhotoUC()
	if res.Error {
		ctx.AbortWithStatusJSON(http.StatusNotFound, res)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
func (h PhotoHandler) UpdatePhoto(ctx *gin.Context) {
	var request baserequest.UpdateRequestPhoto
	err := ctx.ShouldBind(&request)

	if err != nil {
		response := helpers.BuildErrorResponse(err.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	err = validators.ValidateUpdatePhoto(request)
	if err != nil {
		response := helpers.BuildErrorResponse(err.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	res := h.photoUseCase.UpdatePhotoUC(request)
	if res.Error {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
func (h PhotoHandler) DeletePhoto(ctx *gin.Context) {
	id := ctx.Param("photoId")
	res := h.photoUseCase.DeletePhotoUC(id)
	if res.Error {
		ctx.AbortWithStatusJSON(http.StatusNotFound, res)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
