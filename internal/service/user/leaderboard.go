package user

import (
	"context"
	"denet/internal/models/domain"
)

func (s *Service) Leaderboard(ctx context.Context) ([]domain.Leaderboard, error) {
	return s.userRepo.GetLeaderboard(ctx)
}
