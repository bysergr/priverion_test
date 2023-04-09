package utils

import (
	"time"

	"github.com/bysergr/priverion_test/server/dto"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(auth dto.Auth) (string, int64, error) {
	env := GetENV()

	expirationTime := time.Now().Unix() + (60 * 60 * 24 * 12) // 12 days

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":             auth.ID,
		"username":       auth.Username,
		"password":       auth.Password,
		"ExpirationTime": expirationTime,
	})

	tokenString, err := token.SignedString([]byte(env.JWT))
	if err != nil {
		return "", 0, err
	}

	return tokenString, expirationTime, nil
}

func ValidateToken() {

}
