package application

import (
	"api/src/association/domain"
	"api/src/association/domain/ports"
)

type GetAssociationUseCase struct {
	repo ports.AssociationRepository
}

func NewGetAssociationUseCase(repo ports.AssociationRepository) *GetAssociationUseCase {
	return &GetAssociationUseCase{repo: repo}
}

func (useCase *GetAssociationUseCase) Execute(id int) (*domain.Association, error) {
	return useCase.repo.GetByID(id)
}

func (useCase *GetAssociationUseCase) ExecuteAll() ([]domain.Association, error) {
	return useCase.repo.GetAll()
}
