package adapters

import (
	"api/src/user/domain"
	"database/sql"
	"errors"
)

type MySQLUserRepository struct {
	DB *sql.DB
}

func NewMySQLUserRepository(db *sql.DB) *MySQLUserRepository {
	return &MySQLUserRepository{DB: db}
}

func (repo *MySQLUserRepository) CreateUser(user domain.User) (int, error) {
	res, err := repo.DB.Exec("INSERT INTO users (name, email, password, role) VALUES (?, ?, ?, ?)",
		user.Name, user.Email, user.Password, user.Role)

	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), err
}

func (repo *MySQLUserRepository) GetUserByID(id int) (*domain.User, error) {
	var user domain.User
	err := repo.DB.QueryRow("SELECT id, name, email, password, role FROM users WHERE id = ?", id).
		Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}
	
	return &user, err
}

func (repo *MySQLUserRepository) GetAllUsers() ([]domain.User, error) {
	rows, err := repo.DB.Query("SELECT id, name, email, password, role FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var u domain.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.Role); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (repo *MySQLUserRepository) UpdateUser(user domain.User) error {
	_, err := repo.DB.Exec("UPDATE users SET name=?, email=?, password=?, role=? WHERE id=?",
		user.Name, user.Email, user.Password, user.Role, user.ID)
	return err
}

func (repo *MySQLUserRepository) DeleteUser(id int) error {
	res, err := repo.DB.Exec("DELETE FROM users WHERE id=?", id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no se eliminó ningún registro")
	}

	return err
}
