package domain

type Membership struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	AssociationID int  `json:"association_id"`
	Status      string `json:"status"`
	Role        string `json:"role"`
}
