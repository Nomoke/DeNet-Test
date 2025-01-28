package user

import (
	"context"
	"denet/internal/models/domain"
	"errors"
)

func (s *Service) UserStatus(ctx context.Context, userID int) (*domain.UserStatus, error) {
	status, err := s.userRepo.UserStatus(ctx, userID)
	if err != nil {
		return nil, err
	}

	if status == nil {
		return nil, errors.New("user status not found")
	}

	return status, nil
}