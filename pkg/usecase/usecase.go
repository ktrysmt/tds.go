package usecase

import "tds.go/pkg/domain"

type UserUseCase struct {
	userRepo domain.UserRepository
}

func NewUserUseCase(repo domain.UserRepository) *UserUseCase {
	return &UserUseCase{userRepo: repo}
}

func (uc *UserUseCase) CreateUser(name, email, password string) error {
	user := &domain.User{
		Name:     name,
		Email:    email,
		Password: password,
	}
	return uc.userRepo.Save(user)
}

func (uc *UserUseCase) GetUser(id string) (*domain.User, error) {
	return uc.userRepo.FindByID(id)
}
