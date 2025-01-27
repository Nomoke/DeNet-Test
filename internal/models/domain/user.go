package domain

import "time"

type User struct {
	ID       int
	Nickname string
	Password string
}

type UserStatus struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Points         int       `json:"points"`
	RegisteredAt   time.Time `json:"registered_at"`
	TasksCompleted int       `json:"tasks_completed"`
}

type UserCredentials struct {
	Nickname string
	Password string
}
