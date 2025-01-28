package user

type Service struct {
	userRepo UserRepo
}

func NewService(u UserRepo) *Service {
	return &Service{
		userRepo: u,
	}
}
