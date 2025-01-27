package auth

import (
	apperr "denet/internal/pkg/app_err"

	"github.com/golang-jwt/jwt/v5"
)

func (tm *JWTManager) ValidateAccessToken(t string) (*int, error) {
	token, err := jwt.Parse(t, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, apperr.NewUnauthorizedError()
		}

		return []byte(tm.singingKey), nil
	})

	if err != nil {
		return nil, apperr.NewUnauthorizedError()
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, apperr.NewUnauthorizedError()
	}

	userId, ok := claims["sub"].(float64)
	if !ok {
		return nil, apperr.NewUnauthorizedError()
	}

	uId := int(userId)
	return &uId, nil
}
