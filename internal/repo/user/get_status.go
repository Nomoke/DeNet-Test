package user

import (
	"context"
	"denet/internal/models/domain"
)

func (r *Repo) UserStatus(ctx context.Context, userID int) (*domain.UserStatus, error) {
	query := `
		SELECT
			id,
			nickname,
			refferer,
			referals,
			points,
			tasks_completed,
			crated_at
		FROM users
		WHERE id = $1;`

	var status domain.UserStatus

	err := r.pool.QueryRow(ctx, query).Scan(
		&status.ID,
		&status.Nickname,
		&status.Refferer,
		&status.Referals,
		&status.Points,
		&status.TasksCompleted,
		&status.RegisteredAt,
	)
	if err != nil {
		return nil, err
	}

	return &status, nil
}
