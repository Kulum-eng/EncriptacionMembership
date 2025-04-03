package application

import (
	"api/src/user/domain"
	"api/src/user/domain/ports"
)

type GetUserUseCase struct {
	repo ports.UserRepository
}

func NewGetUserUseCase(repo ports.UserRepository) *GetUserUseCase {
	return &GetUserUseCase{repo: repo}
}

func (uc *GetUserUseCase) ExecuteByID(id int) (*domain.User, error) {
	return uc.repo.GetUserByID(id)
}

func (uc *GetUserUseCase) ExecuteAll() ([]domain.User, error) {
	return uc.repo.GetAllUsers()
}
