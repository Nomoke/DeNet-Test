package task

import (
	"context"
	"denet/internal/models/domain"
	apperr "denet/internal/pkg/app_err"
)

func (r *Service) TaskComplete(ctx context.Context, ut domain.UserTaskComplete) error {
	isCompletes, err :=  r.TaskRepo.IsTaskDone(ctx, ut)
	if err != nil {
		return err
	}

	if isCompletes {
		return apperr.NewAuthorizationError("Вы уже выполнили эту задачу")
	}


	return r.TaskRepo.TaskComplete(ctx, ut)
}