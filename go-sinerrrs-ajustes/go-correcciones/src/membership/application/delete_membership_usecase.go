package application

import (
	"api/src/membership/domain/ports"
	"errors"
)

type DeleteMembershipUseCase struct {
	repo ports.MembershipRepository
}

func NewDeleteMembershipUseCase(repo ports.MembershipRepository) *DeleteMembershipUseCase {
	return &DeleteMembershipUseCase{repo: repo}
}

func (uc *DeleteMembershipUseCase) Execute(id int) error {
	membership, err := uc.repo.GetMembershipByID(id)
	if err != nil {
		return err
	}

	if membership == nil {
		return errors.New("el miembro no existe")
	}

	return uc.repo.DeleteMembership(id)
}
