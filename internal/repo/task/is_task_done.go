package task

import (
	"context"
	"denet/internal/models/domain"

	"github.com/pkg/errors"
)

func (r *Repo) IsTaskDone(ctx context.Context, ut domain.UserTaskComplete) (bool, error) {
	var exists bool
	query := `
        SELECT EXISTS (
            SELECT 1 
            FROM user_tasks 
            WHERE user_id = $1 AND task_id = $2
        );`

	err := r.pool.QueryRow(ctx, query, ut.UserID, ut.TaskID).Scan(&exists)
	if err != nil {
		return false, errors.WithStack(err)
	}

	return exists, nil
}
