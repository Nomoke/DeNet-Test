package task

import (
	"context"
	"denet/internal/models/domain"

	"github.com/pkg/errors"
)

func (r *Repo) EarnUserPoints(ctx context.Context, ut domain.UserTaskComplete) error {

	return nil
}

func (r *Repo) getTaskPoints(ctx context.Context, taskID int) (int, error) {
	var points int
    query := `
        SELECT 
			points
        FROM tasks
        WHERE id = $1;
    `

	err := r.pool.QueryRow(ctx, query, taskID).Scan(&points)
	if err != nil {
		return 0, errors.WithStack(err)
	}

	return points, nil
}
