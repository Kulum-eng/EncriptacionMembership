package ports

import "api/src/user/domain"

type UserRepository interface {
	CreateUser(user domain.User) (int, error)
	GetUserByID(id int) (*domain.User, error)
	GetAllUsers() ([]domain.User, error)
	UpdateUser(user domain.User) error
	DeleteUser(id int) error
}
