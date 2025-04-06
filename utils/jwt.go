package utils

import (
	"fmt"
	"time"
	"tripat3k2/url_shortner/config"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	UserId uint `json:"userId"`
	jwt.MapClaims
}

func CreateToken(userId uint, multiplier int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		UserId: userId,
		MapClaims: jwt.MapClaims{
			"exp": time.Now().Add(time.Minute * time.Duration(multiplier)).Unix(),
			"iat": time.Now().Unix(),
		},
	})

	tokenString, err := token.SignedString([]byte(config.Env.JWT_SECRET))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (*Claims, error) {
	claims := new(Claims)
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(config.Env.JWT_SECRET), nil
	})

	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid or expired token")
	}

	return claims, nil
}
