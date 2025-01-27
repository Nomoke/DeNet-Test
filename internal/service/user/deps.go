package user

import (
	"context"

	"denet/internal/models/domain"
)

type UserRepo interface {
	GetLeaderboard(ctx context.Context) ([]domain.Leaderboard, error)
}
