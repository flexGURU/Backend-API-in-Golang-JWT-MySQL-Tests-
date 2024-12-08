package auth

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pwd string) (string, error) {
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	
	if err != nil {
		return "", err
	}

	return string(hashedPwd), nil
}


func ComparePassword(hashedpwd string, pwd []byte) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hashedpwd), pwd)

	return err == nil




}