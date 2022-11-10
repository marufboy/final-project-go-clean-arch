package http

import (
	"final-project-go-clean-arch/domain"
	"final-project-go-clean-arch/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SocialMediaHandler struct {
	socialMediaUseCase domain.SocialMediaUseCase
}

func NewSocialMediaHandler(r *gin.Engine, socialMediaUC domain.SocialMediaUseCase, db *gorm.DB) {
	handler := &SocialMediaHandler{
		socialMediaUseCase: socialMediaUC,
	}

	router := r.Group("/socialmedias", middlewares.AuthorizeJWT())
	{
		router.POST("/", handler.CreateSocialMedia)
		router.GET("/", handler.GetSocialMedia)
		router.PUT("/:socialMediaId", middlewares.AuthorizationSocialMedia(db), handler.UpdateSocialMedia)
		router.DELETE("/:socialMediaId", middlewares.AuthorizationSocialMedia(db), handler.DeleteSocialMedia)
	}
}

func (h SocialMediaHandler) CreateSocialMedia(ctx *gin.Context) {}
func (h SocialMediaHandler) GetSocialMedia(ctx *gin.Context)    {}
func (h SocialMediaHandler) UpdateSocialMedia(ctx *gin.Context) {}
func (h SocialMediaHandler) DeleteSocialMedia(ctx *gin.Context) {}
