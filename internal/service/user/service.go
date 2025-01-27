package user

import (
	"context"
	"denet/internal/models/domain"
)

type Service struct {
	userRepo UserRepo
}

func (s *Service) Leaderboard(ctx context.Context) ([]domain.Leaderboard, error) {
	return s.userRepo.GetLeaderboard(ctx)
}
