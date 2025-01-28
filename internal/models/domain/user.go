package domain

import "time"

type User struct {
	ID       int
	Nickname string
	Password string
}

type UserStatus struct {
	ID             int
	Nickname       string
	Refferer       string
	Referals       int
	Points         int
	TasksCompleted int
	RegisteredAt   time.Time
}

type UserCredentials struct {
	Nickname string
	Password string
}
