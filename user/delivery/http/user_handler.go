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

type UserHandler struct {
	userUseCase domain.UserUseCase
}

func NewUserHandler(r *gin.Engine, usecase domain.UserUseCase, db *gorm.DB) {
	handler := &UserHandler{
		userUseCase: usecase,
	}

	router := r.Group("/user")
	router.POST("/register", handler.handleRegister)
	router.POST("/login", handler.handleLogin)
	{
		router.Use(middlewares.AuthorizeJWT())
		router.PUT("/:userId", handler.handleUpdateUser)
		router.DELETE("/:userId", handler.handleDeleteUser)
	}
}

func (h UserHandler) handleRegister(ctx *gin.Context) {
	var request baserequest.RegisterRequest
	err := ctx.ShouldBind(&request)

	if err != nil {
		response := helpers.BuildErrorResponse(err.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	//handling validate struct
	err = validators.ValidateRegister(request)
	if err != nil {
		response := helpers.BuildErrorResponse(err.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	//TODO: usecase register
	res := h.userUseCase.UserRegisterUC(request)
	if res.Error {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
func (h UserHandler) handleLogin(ctx *gin.Context) {
	var request baserequest.LoginRequest
	err := ctx.ShouldBind(&request)

	if err != nil {
		response := helpers.BuildErrorResponse(err.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	//handling validate struct
	err = validators.ValidateLogin(request)
	if err != nil {
		response := helpers.BuildErrorResponse(err.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	res := h.userUseCase.UserLoginUC(request)
	if res.Error {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
func (h UserHandler) handleUpdateUser(ctx *gin.Context) {
	var request baserequest.UpdateRequestUser
	err := ctx.ShouldBind(&request)

	if err != nil {
		response := helpers.BuildErrorResponse(err.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	//handling validate struct
	err = validators.ValidateUpdateUser(request)
	if err != nil {
		response := helpers.BuildErrorResponse(err.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	res := h.userUseCase.UpdateUserUC(request)
	if res.Error {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	ctx.JSON(http.StatusCreated, res)
}
func (h UserHandler) handleDeleteUser(ctx *gin.Context) {
	id := ctx.Param("userId")
	res := h.userUseCase.DeleteUserUC(id)
	if res.Error {
		ctx.AbortWithStatusJSON(http.StatusNotFound, res)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
