package domain

type UserTaskComplete struct {
	UserID int
	TaskID int
}

type Task struct {
	ID     int
	Name   string
	Points int
	Desctiption string
}