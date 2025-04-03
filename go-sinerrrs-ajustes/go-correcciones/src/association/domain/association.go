package domain

type Association struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	Contact     string `json:"contact"`
	Services    string `json:"services"`
}
