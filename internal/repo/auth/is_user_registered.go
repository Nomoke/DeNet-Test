package auth

import (
	"context"
	"database/sql"
)

func (r *Repo) IsUserRegistered(ctx context.Context, nickname string) (int, error) {
	var uID int
	query := `
		SELECT
		id FROM users 
		WHERE nickname = $1 
		LIMIT 1;`

	err := r.pool.QueryRow(ctx, query, nickname).Scan(&uID)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, err
	}

	return uID, nil
}
