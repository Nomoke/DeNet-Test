package models

type UserStatusResp struct {
	Nickname       string `json:"nickname"`
	Points         int    `json:"points"`
	TasksCompleted int    `json:"tasksCompleted"`
	RegisteredAt   string `json:"registeredAt"`
}
