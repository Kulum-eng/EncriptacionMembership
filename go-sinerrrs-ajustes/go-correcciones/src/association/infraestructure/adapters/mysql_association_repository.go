package adapters

import (
	"api/src/association/domain"
	"database/sql"
)

type MySQLAssociationRepository struct {
	db *sql.DB
}

func NewMySQLAssociationRepository(db *sql.DB) *MySQLAssociationRepository {
	return &MySQLAssociationRepository{db: db}
}

func (r *MySQLAssociationRepository) Create(association domain.Association) (int, error) {
	query := "INSERT INTO associations (name, address, contact, services) VALUES (?, ?, ?, ?)"
	
	res, err := r.db.Exec(query, association.Name, association.Contact, association.Contact, association.Services)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	
	return int(id), err
}

func (r *MySQLAssociationRepository) GetByID(id int) (*domain.Association, error) {
	var a domain.Association
	
	query := "SELECT id, name, address, contact, services FROM associations WHERE id = ?"
	err := r.db.QueryRow(query, id).Scan(
		&a.ID,
		&a.Name,
		&a.Contact,
		&a.Services,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &a, nil
}


func (repo *MySQLAssociationRepository) GetAll() ([]domain.Association, error) {
	query := "SELECT id, name, address, contact, services FROM associations"
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var associations []domain.Association
	for rows.Next() {
		var a domain.Association
		if err := rows.Scan(&a.ID, &a.Name, &a.Contact, &a.Services); err != nil {
			return nil, err
		}
		associations = append(associations, a)
	}

	return associations, nil
}

func (r *MySQLAssociationRepository) Update(association domain.Association) error {
	query := "UPDATE associations SET name = ?, address = ?, contact = ?, services = ? WHERE id = ?"
	_, err := r.db.Exec(query, association.Name, association.Address, association.Contact, association.Services, association.ID)
	return err
}

func (r *MySQLAssociationRepository) Delete(id int) error {
	query := "DELETE FROM associations WHERE id = ?"
	_, err := r.db.Exec(query, id)
	return err
}
