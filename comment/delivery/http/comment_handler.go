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

type CommentHandler struct {
	commentUseCase domain.CommentUseCase
}

func NewCommentHandler(r *gin.Engine, commentUC domain.CommentUseCase, db *gorm.DB) {
	handler := &CommentHandler{
		commentUseCase: commentUC,
	}

	router := r.Group("/comments", middlewares.AuthorizeJWT())
	{
		router.POST("/", handler.CreateComment)
		router.GET("/", handler.GetComment)
		router.PUT("/:commentId", middlewares.AuthorizationComment(db), handler.UpdateComment)
		router.DELETE("/:commentId", middlewares.AuthorizationComment(db), handler.DeleteComment)
	}
}

func (h CommentHandler) CreateComment(ctx *gin.Context) {
	var request baserequest.CreateRequestComment
	err := ctx.ShouldBind(&request)

	if err != nil {
		response := helpers.BuildErrorResponse(err.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	err = validators.ValidateCreateComment(request)
	if err != nil {
		response := helpers.BuildErrorResponse(err.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	res := h.commentUseCase.CreateCommentUC(request)
	if res.Error {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
func (h CommentHandler) GetComment(ctx *gin.Context) {
	res := h.commentUseCase.GetCommentUC()
	if res.Error {
		ctx.AbortWithStatusJSON(http.StatusNotFound, res)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
func (h CommentHandler) UpdateComment(ctx *gin.Context) {
	var request baserequest.UpdateRequestComment
	err := ctx.ShouldBind(&request)

	if err != nil {
		response := helpers.BuildErrorResponse(err.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	err = validators.ValidateUpdateComment(request)
	if err != nil {
		response := helpers.BuildErrorResponse(err.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	res := h.commentUseCase.UpdateCommentUC(request)
	if res.Error {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
func (h CommentHandler) DeleteComment(ctx *gin.Context) {
	id := ctx.Param("commentId")
	res := h.commentUseCase.DeleteCommentUC(id)
	if res.Error {
		ctx.AbortWithStatusJSON(http.StatusNotFound, res)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
