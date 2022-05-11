package helper

import (
	"sejutacita/config"
	"time"

	"github.com/golang-jwt/jwt"
	// "github.com/dgrijalva/jwt-go"
)

func TokenJWT(role string, email string, username string) (string, string, string) {
	token := jwt.New(jwt.SigningMethodHS256)
	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["role"] = role
	claims["email"] = email
	claims["username"] = username
	//claims["exp"] = time.Now().Add(time.Second * 30).Unix()
	//claims["exp"] = time.Now().Add(time.Minute * 1).Unix()
	claims["exp"] = time.Now().Add(time.Hour * 3).Unix()

	// Generate encoded token and send it as response.
	restoken, err := token.SignedString([]byte(config.AppKey))
	if err != nil {
		return "", "81", err.Error()
	}
	return restoken, "00", "Success"
}
