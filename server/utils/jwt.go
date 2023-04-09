package utils

import (
	"errors"
	"fmt"
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
		"ExpirationTime": expirationTime,
	})

	tokenString, err := token.SignedString([]byte(env.JWT))
	if err != nil {
		return "", 0, err
	}

	return tokenString, expirationTime, nil
}

func ValidateToken(tokenString string)  (map[string]string, error) {
	env := GetENV()

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("error validating token")
		}


		return []byte(env.JWT), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return map[string]string{
			"id":             fmt.Sprintf("%v", claims["id"]),
			"username":       fmt.Sprintf("%v", claims["username"]),
			"ExpirationTime": fmt.Sprintf("%v", claims["ExpirationTime"]),
		}, nil
	}

	return nil, errors.New("error validating token")
}
