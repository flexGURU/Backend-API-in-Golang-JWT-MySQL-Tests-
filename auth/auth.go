package auth

import (
	"fmt"
	"log"
	"net/http"

	"github.com/flexGURU/goAPI/config"
	"github.com/flexGURU/goAPI/types"
	"github.com/flexGURU/goAPI/utils"
	"github.com/golang-jwt/jwt/v5"
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

func WithJWTAuth(handler http.HandlerFunc, store types.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get Token From User request
		tokenAuthStr := getToken(r)

		// Validate the token
		token, err := validateToken(tokenAuthStr)
		if err != nil {
			log.Printf("failed to verify token %v", err)
			permissonDenied(w)
			return
		}

		if !token.Valid {
			log.Printf("token is not valid")
			permissonDenied(w)
			return
		}




	}
}

// Helper function to get the token
func getToken(r *http.Request) string {

	tokenAuth := r.Header.Get("Authorization")

	if tokenAuth == "" {
		return "token auth is missing"

	}

	return tokenAuth

}

// helper function to validate the token
func validateToken(t string) (*jwt.Token, error) {
	return jwt.Parse(t, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t.Header["alg"])
		}
		
		return []byte(config.Envs.JWTSecret), nil
	})
}

func permissonDenied(w http.ResponseWriter) {
	utils.WriteError(w, http.StatusForbidden, fmt.Errorf("persmission denied"))
}