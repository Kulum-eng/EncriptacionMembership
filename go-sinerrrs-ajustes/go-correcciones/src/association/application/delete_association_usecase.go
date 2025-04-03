package application

import (
	"api/src/association/domain/ports"
	"errors"
)

type DeleteAssociationUseCase struct {
	repo ports.AssociationRepository
}

func NewDeleteAssociationUseCase(repo ports.AssociationRepository) *DeleteAssociationUseCase {
	return &DeleteAssociationUseCase{repo: repo}
}

func (useCase *DeleteAssociationUseCase) Execute(id int) error {
	association, err := useCase.repo.GetByID(id)
	if err != nil {
		return err
	}

	if association == nil {
		return errors.New("la asociación no existe")
	}

	return useCase.repo.Delete(id)
}
