package application

import (
	"api/src/user/domain/ports"
	"errors"
)

type DeleteUserUseCase struct {
	repo ports.UserRepository
}

func NewDeleteUserUseCase(repo ports.UserRepository) *DeleteUserUseCase {
	return &DeleteUserUseCase{repo: repo}
}

func (uc *DeleteUserUseCase) Execute(id int) error {
	user, err := uc.repo.GetUserByID(id)
	if err != nil {
		return err
	}

	if user == nil {
		return errors.New("el usuario no existe")
	}

	return uc.repo.DeleteUser(id)
}
