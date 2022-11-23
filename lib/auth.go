package lib

import (
	"log"
	"os"

	"github.com/golang-jwt/jwt"
)

var AUTH_PROVIDER_TWITTER int = 0
var AUTH_PROVIDER_MASTODON int = 1

func ValidateToken(tokenStr string) (jwt.MapClaims, bool) {
	hmacSecretString := os.Getenv("JWT_SECRET")
	hmacSecret := []byte(hmacSecretString)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// check token signing method etc
		return hmacSecret, nil
	})

	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		log.Printf("Invalid JWT Token")
		return nil, false
	}
}
