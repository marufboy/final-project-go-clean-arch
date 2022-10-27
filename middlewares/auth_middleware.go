package middlewares

import (
	"final-project-go-clean-arch/domain"
	"final-project-go-clean-arch/helpers"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			res := helpers.BuildErrorResponse("Failed to process request", nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}

		bearer := strings.HasPrefix(authHeader, "Bearer")
		if !bearer {
			res := helpers.BuildErrorResponse("Failed to process request", nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}

		stringToken := strings.Split(authHeader, " ")[1]
		token := helpers.VerifyToken(stringToken, ctx)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			ctx.Set("user", claims)
			ctx.Next()
		} else {
			res := helpers.BuildErrorResponse("Token is not valid", nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
		}
	}
}

func AuthorizationUser(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		param, err := strconv.Atoi(ctx.Param("userId"))
		if err != nil {
			res := helpers.BuildErrorResponse("invalid parameter", nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}

		userData := ctx.MustGet("user").(jwt.MapClaims)
		userId := uint(userData["id"].(float64))

		User := domain.User{}
		err = db.Select("id").First(&User, uint(param)).Error
		if err != nil {
			res := helpers.BuildErrorResponse("user not found", nil)
			ctx.AbortWithStatusJSON(http.StatusNotFound, res)
			return
		}

		if User.ID != userId {
			res := helpers.BuildErrorResponse("not allowed to access the data", nil)
			ctx.AbortWithStatusJSON(http.StatusForbidden, res)
			return
		}

		ctx.Next()
	}
}

func AuthorizationPhoto(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		param, err := strconv.Atoi(ctx.Param("photoId"))
		if err != nil {
			res := helpers.BuildErrorResponse("invalid parameter", nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}

		userData := ctx.MustGet("user").(jwt.MapClaims)
		userId := uint(userData["id"].(float64))

		Photo := domain.Photo{}
		err = db.Select("id").First(&Photo, uint(param)).Error
		if err != nil {
			res := helpers.BuildErrorResponse("user not found", nil)
			ctx.AbortWithStatusJSON(http.StatusNotFound, res)
			return
		}

		if Photo.UserID != userId {
			res := helpers.BuildErrorResponse("not allowed to access the data", nil)
			ctx.AbortWithStatusJSON(http.StatusForbidden, res)
			return
		}

		ctx.Next()
	}
}

func AuthorizationComment(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		param, err := strconv.Atoi(ctx.Param("commentId"))
		if err != nil {
			res := helpers.BuildErrorResponse("invalid parameter", nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}

		userData := ctx.MustGet("user").(jwt.MapClaims)
		userId := uint(userData["id"].(float64))

		Comment := domain.Comment{}
		err = db.Select("id").First(&Comment, uint(param)).Error
		if err != nil {
			res := helpers.BuildErrorResponse("user not found", nil)
			ctx.AbortWithStatusJSON(http.StatusNotFound, res)
			return
		}

		if Comment.UserID != userId {
			res := helpers.BuildErrorResponse("not allowed to access the data", nil)
			ctx.AbortWithStatusJSON(http.StatusForbidden, res)
			return
		}

		ctx.Next()
	}
}

func AuthorizationSocialMedia(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		param, err := strconv.Atoi(ctx.Param("socialMediaId"))
		if err != nil {
			res := helpers.BuildErrorResponse("invalid parameter", nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}

		userData := ctx.MustGet("user").(jwt.MapClaims)
		userId := uint(userData["id"].(float64))

		SocialMedia := domain.SocialMedia{}
		err = db.Select("id").First(&SocialMedia, uint(param)).Error
		if err != nil {
			res := helpers.BuildErrorResponse("user not found", nil)
			ctx.AbortWithStatusJSON(http.StatusNotFound, res)
			return
		}

		if SocialMedia.UserID != userId {
			res := helpers.BuildErrorResponse("not allowed to access the data", nil)
			ctx.AbortWithStatusJSON(http.StatusForbidden, res)
			return
		}

		ctx.Next()
	}
}
