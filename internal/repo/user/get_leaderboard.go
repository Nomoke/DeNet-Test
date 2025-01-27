package user

import (
	"context"

	"denet/internal/models/domain"
)

func (r *Repo) GetLeaderboard(ctx context.Context) ([]domain.Leaderboard, error) {
	query := `
		SELECT
			id, 
			points,
			nickname
		FROM users
		ORDER BY points DESC
		LIMIT 10;`

	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var lbList []domain.Leaderboard

	for rows.Next() {
		var lb domain.Leaderboard
		if err := rows.Scan(&lb.ID, &lb.Points, &lb.Nickname); err != nil {
			return nil, err
		}

		lbList = append(lbList, lb)
	}

	return lbList, nil
}
