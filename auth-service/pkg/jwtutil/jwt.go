package jwtutil

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const tokenTTL = 72 * time.Hour

func Generate(secret, userID, email string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"exp":     time.Now().Add(tokenTTL).Unix(),
		"iat":     time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
