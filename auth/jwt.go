package auth

import (
	"strconv"
	"time"

	"github.com/flexGURU/goAPI/config"
	"github.com/golang-jwt/jwt/v5"
)


func CreateJWT(secret []byte, userID int64) (string, error) {

	expiryDuration := time.Second * time.Duration(
		config.
	)

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userID" : strconv.Itoa(userID),
			"expiredAt" : 
		}
	)

}