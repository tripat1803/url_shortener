package utils

import (
	"time"
	"tripat3k2/url_shortner/config"

	"github.com/golang-jwt/jwt"
)

func CreateToken(userId uint, multiplier int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(time.Minute * time.Duration(multiplier)).Unix(),
		"iat":    time.Now().Unix(),
	})

	tokenString, err := token.SignedString([]byte(config.Env.JWT_SECRET))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
