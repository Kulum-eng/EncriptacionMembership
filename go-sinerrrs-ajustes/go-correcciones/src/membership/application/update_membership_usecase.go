package application

import (
	"api/src/membership/domain"
	"api/src/membership/domain/ports"
	"errors"
)

type UpdateMembershipUseCase struct {
	repo ports.MembershipRepository
}
func NewUpdateMembershipUseCase(repo ports.MembershipRepository) *UpdateMembershipUseCase {
	return &UpdateMembershipUseCase{repo: repo}
}

func (uc *UpdateMembershipUseCase) Execute(membership domain.Membership) error {
	membershipToUpdate, err := uc.repo.GetMembershipByID(membership.ID)
	if err != nil {
		return err
	}

	if membershipToUpdate == nil {
		return errors.New("el miembro no existe")
	}
	
	return uc.repo.UpdateMembership(membership)
}
