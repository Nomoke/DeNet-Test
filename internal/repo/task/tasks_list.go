package task

import (
	"context"

	"denet/internal/models/domain"

	"github.com/pkg/errors"
)

func (r *Repo) List(ctx context.Context) ([]domain.Task, error) {
	query := `
		SELECT
			id,
			name,
			points,
			description
		FROM tasks
		WHERE is_active = TRUE;`

	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	defer rows.Close()

	var tasks []domain.Task

	for rows.Next() {
		var t domain.Task
		if err := rows.Scan(
			&t.ID,
			&t.Name,
			&t.Points,
			&t.Desctiption,
		); err != nil {
			return nil, errors.WithStack(err)
		}

		tasks = append(tasks, t)
	}

	return tasks, nil
}
