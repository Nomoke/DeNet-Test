package task

import (
	"context"

	"denet/internal/models/domain"
)

type TaskRepo interface {
	List(ctx context.Context) ([]domain.Task, error)
	IsTaskDone(ctx context.Context, ut domain.UserTaskComplete) (bool, error)
	TaskComplete(ctx context.Context, ut domain.UserTaskComplete) error
}
