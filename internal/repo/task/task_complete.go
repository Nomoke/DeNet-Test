package task

import (
	"context"
	"denet/internal/models/domain"
)

func (r *Repo) TaskComplete(ctx context.Context, ut domain.UserTaskComplete) error {
	query := `
		INSERT INTO user_tasks (user_id, task_id)
		VALUES ($1, $2)`

	_, err := r.pool.Exec(ctx, query, ut.UserID, ut.TaskID)
	if err != nil {
		return err
	}

	return nil
}
