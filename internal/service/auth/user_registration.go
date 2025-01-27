package auth

import (
	"context"
	"denet/internal/models/domain"
	apperr "denet/internal/pkg/app_err"
)

func (s *Service) UserRegistration(ctx context.Context, userReg domain.UserCredentials) (*domain.Tokens, error) {
	id, err := s.authRepo.IsUserRegistered(ctx, userReg.Nickname)
	if err != nil {
		return nil, err
	}

	if id > 0 {
		return nil, apperr.NewAuthorizationError("Пользователь с таким логином уже cуществует")
	}

	uID, err := s.authRepo.CreateUser(ctx, &userReg)
	if err != nil {
		return nil, err
	}

	accessToken, err := s.jwtRepo.AccessToken(uID)
	if err != nil {
		return nil, err
	}

	tokens := domain.Tokens{
		Access: accessToken,
	}

	return &tokens, nil
}
