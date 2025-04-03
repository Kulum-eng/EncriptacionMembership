package application

import (
	"api/src/user/domain"
	"api/src/user/domain/ports"
	"errors"
)

type UpdateUserUseCase struct {
	repo ports.UserRepository
}

func NewUpdateUserUseCase(repo ports.UserRepository) *UpdateUserUseCase {
	return &UpdateUserUseCase{repo: repo}
}

func (uc *UpdateUserUseCase) Execute(user domain.User) error {
	userToUpdate, err := uc.repo.GetUserByID(user.ID)
	if err != nil {
		return err
	}

	if userToUpdate == nil {
		return errors.New("el usuario no existe")
	}
	
	return uc.repo.UpdateUser(user)
}
