package auth

import (
	"context"
	"denet/internal/models/domain"

	"github.com/pkg/errors"
)

func (r *Repo) CreateUser(ctx context.Context, userReg domain.UserCredentials) (*int, error) {
	var uID int
	query := `
		INSERT INTO users (email, password)
		VALUES ($1, $2)
		RETURNING id`

	err := r.pool.QueryRow(ctx, query, userReg.Nickname, userReg.Password).Scan(&uID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &uID, nil
}
