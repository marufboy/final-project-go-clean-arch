package helpers

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error load .env file")
	}
}

var SECRET_KEY = os.Getenv("SECRET_KEY")

type jwtCustomClaim struct {
	UserID     uint `json:"user_id"`
	Authorized bool `json:"authorized"`
	jwt.RegisteredClaims
}

func GenerateToken(UserID uint) string {
	claims := &jwtCustomClaim{
		UserID,
		true,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().AddDate(0, 0, 1)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	signedToken, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Println(err)
		log.Fatal("Failed to sign token")
	}

	return signedToken
}

func VerifyToken(token string, ctx *gin.Context) *jwt.Token {
	verify, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("token invalid, request a new one")
		}
		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return nil
	}

	return verify
}
