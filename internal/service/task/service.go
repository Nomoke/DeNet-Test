package task

type Service struct {
	TaskRepo TaskRepo
}

func New(t TaskRepo) *Service {
	return &Service{
		TaskRepo: t}
}
