package helpers

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(pwd string) string {
	password := []byte(pwd)
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		log.Fatal("Failed to hash a password")
	}

	return string(hash)
}

func ComparePass(hashedPwd string, pwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, pwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
