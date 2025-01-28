package task

import (
	"context"
	"denet/internal/models/domain"
)

func(s *Service) List(ctx context.Context) ([]domain.Task, error) {
	return s.TaskRepo.List(ctx)
}