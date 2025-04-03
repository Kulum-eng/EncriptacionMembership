package application

import (
	"api/src/association/domain"
	"api/src/association/domain/ports"
)

type CreateAssociationUseCase struct {
	repo ports.AssociationRepository
}

func NewCreateAssociationUseCase(repo ports.AssociationRepository) *CreateAssociationUseCase {
	return &CreateAssociationUseCase{repo: repo}
}

func (useCase *CreateAssociationUseCase) Execute(association domain.Association) (int, error) {
	return useCase.repo.Create(association)
}
