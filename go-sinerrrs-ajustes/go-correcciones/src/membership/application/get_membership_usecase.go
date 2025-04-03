package application

import (
	"api/src/membership/domain"
	"api/src/membership/domain/ports"
)

type GetMembershipUseCase struct {
	repo ports.MembershipRepository
}

func NewGetMembershipUseCase(repo ports.MembershipRepository) *GetMembershipUseCase {
	return &GetMembershipUseCase{repo: repo}
}

func (uc *GetMembershipUseCase) ExecuteByID(id int) (*domain.Membership, error) {
	return uc.repo.GetMembershipByID(id)
}

func (uc *GetMembershipUseCase) ExecuteAll() ([]domain.Membership, error) {
	return uc.repo.GetAllMemberships()
}
