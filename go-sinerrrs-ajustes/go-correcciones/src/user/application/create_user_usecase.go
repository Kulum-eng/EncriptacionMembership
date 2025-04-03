package application

import (
	"api/src/user/domain"
	"api/src/user/domain/ports"
)

type CreateUserUseCase struct {
	repo ports.UserRepository
}

func NewCreateUserUseCase(repo ports.UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{repo: repo}
}

func (uc *CreateUserUseCase) Execute(user domain.User) (int, error) {
	return uc.repo.CreateUser(user)
}
