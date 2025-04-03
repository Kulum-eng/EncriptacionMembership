package responses

import "api/src/membership/domain"

type MembershipResponse struct {
	ID            int    `json:"id"`
	UserID        int    `json:"user_id"`
	AssociationID int    `json:"association_id"`
	Status        string `json:"status"`
	Role          string `json:"role"`
}

func NewMembershipResponse(m domain.Membership) MembershipResponse {
	return MembershipResponse{
		ID:            m.ID,
		UserID:        m.UserID,
		AssociationID: m.AssociationID,
		Status:        m.Status,
		Role:          m.Role,
	}
}

func NewMembershipListResponse(memberships []domain.Membership) []MembershipResponse {
	var responses []MembershipResponse
	for _, m := range memberships {
		responses = append(responses, NewMembershipResponse(m))
	}
	return responses
}
