package application

import (
	"api/src/membership/domain"
	"api/src/membership/domain/ports"
)

type CreateMembershipUseCase struct {
	repo ports.MembershipRepository
}

func NewCreateMembershipUseCase(repo ports.MembershipRepository) *CreateMembershipUseCase {
	return &CreateMembershipUseCase{repo: repo}
}

func (uc *CreateMembershipUseCase) Execute(membership domain.Membership) (int, error) {
	return uc.repo.CreateMembership(membership)
}
