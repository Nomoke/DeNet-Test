package auth

import (
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

type JWTManager struct {
	singingKey string
}

func NewJWTManager(singingKey string) *JWTManager {
	return &JWTManager{singingKey: singingKey}
}

func (tm *JWTManager) AccessToken(userId int) (string, error) {
	claims := jwt.MapClaims{
		"sub": userId,
		"exp": time.Now().Add(15 * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(tm.singingKey))
}
