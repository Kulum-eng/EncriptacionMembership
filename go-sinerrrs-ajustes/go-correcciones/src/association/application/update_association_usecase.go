package application

import (
	"api/src/association/domain"
	"api/src/association/domain/ports"
	"errors"
)

type UpdateAssociationUseCase struct {
	repo ports.AssociationRepository
}

func NewUpdateAssociationUseCase(repo ports.AssociationRepository) *UpdateAssociationUseCase {
	return &UpdateAssociationUseCase{repo: repo}
}

func (useCase *UpdateAssociationUseCase) Execute(association domain.Association) error {
	associationToUpdate, err := useCase.repo.GetByID(association.ID)
	if err != nil {
		return err
	}

	if associationToUpdate == nil {
		return errors.New("la asociación no existe")
	}

	return useCase.repo.Update(association)
}
