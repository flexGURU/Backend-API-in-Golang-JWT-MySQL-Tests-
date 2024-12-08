package auth

import (
	"strconv"
	"time"

	"github.com/flexGURU/goAPI/config"
	"github.com/golang-jwt/jwt/v5"
)


func CreateJWT(secret []byte, userID int) (string, error) {

	expiryDuration := time.Second * time.Duration(config.Envs.JWTDuration)

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userID" : strconv.Itoa(userID),
			"expiredAt" : time.Now().Add(expiryDuration).Unix(),
		},
	)

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "DDD", err
	}

	return tokenString, nil

}