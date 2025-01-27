package domain

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Session struct {
	ID           string
	UserID       int
	RefreshToken string
	ExpiredAt    time.Time
}

type Tokens struct {
	Access string
}

type Claims struct {
	jwt.RegisteredClaims
	UserID    string
	SessionID string
}
