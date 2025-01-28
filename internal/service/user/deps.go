package user

import (
	"context"

	"denet/internal/models/domain"
)

type UserRepo interface {
	UserStatus(ctx context.Context, userID int) (*domain.UserStatus, error)
	GetLeaderboard(ctx context.Context) ([]domain.Leaderboard, error)
}
