package models

type Task struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Points      int    `json:"points"`
	Description string `json:"description"`
}

type TaskCompleteReq struct {
	TaskID int `json:"taskId" validate:"required,gte=1"`
}
