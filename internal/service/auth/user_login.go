package auth

import (
	"context"
	"denet/internal/models/domain"
	apperr "denet/internal/pkg/app_err"
)

func (s *Service) UserLogin(ctx context.Context, user domain.UserCredentials) (*domain.Tokens, error) {
	uID, err := s.authRepo.IsUserRegistered(ctx, user.Nickname)
	if err != nil {
		return nil, err
	}

	if uID == 0 {
		return nil, apperr.NewAuthorizationError("Пользователь с таким логином не зарегистрирован")
	}

	token, err := s.jwtRepo.AccessToken(uID)
	if err != nil {
		return nil, err
	}

	tokens := domain.Tokens{
		Access: token,
	}

	return &tokens, nil
}
