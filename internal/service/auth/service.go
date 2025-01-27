package auth

type Service struct {
	authRepo authRepo
	jwtRepo  jwtRepo
}

func NewService(a authRepo, j jwtRepo) *Service {
	return &Service{
		authRepo: a,
		jwtRepo:  j,
	}
}
