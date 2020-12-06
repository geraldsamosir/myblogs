package helper

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type Password struct{}

func (*Password) HashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)

	if err != nil {
		panic("Failed to generate hash password")
	}

	return string(hash)
}

func (*Password) ComparePassword(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)

	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		logrus.Error("Error in compare password")
		return false
	}
	return true
}
