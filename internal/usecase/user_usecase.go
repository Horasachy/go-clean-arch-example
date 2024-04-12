package usecase

type UserUseCase struct {
	userRepo *UserRepository
}

func (uuc *UserUseCase) Create() {
	uuc.userRepo
}
