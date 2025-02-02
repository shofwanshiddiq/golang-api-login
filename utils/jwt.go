package utils

import (
	"api-integration/config"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(userId uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(config.GetJWTExpirationTime()).Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.GetJWTSecret())
}

func ValidateToken(token string) (uint, error) {
	tokens, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return config.GetJWTSecret(), nil
	})
	if err != nil {
		return 0, err
	}
	if claims, ok := tokens.Claims.(jwt.MapClaims); ok && tokens.Valid {
		userId := uint(claims["user_id"].(float64))
		return userId, nil
	}

	return 0, nil
}
