package auth

import (
	"context"

	"denet/internal/models/domain"
)

type authRepo interface {
	CreateUser(ctx context.Context, user *domain.UserCredentials) (int, error)
	UserLogin(email string, password string) (*domain.Tokens, error)
	IsUserRegistered(ctx context.Context, nickname string) (int, error)
}

type jwtRepo interface {
	AccessToken(userId int) (string, error)
	ValidateAccessToken(token string) (*int, error)
}
