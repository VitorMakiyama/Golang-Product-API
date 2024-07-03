package repositories

import (
	"api-produtos/internal/core/domain"
	"api-produtos/internal/core/ports"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

type mySQLTypeRepositoryAdapter struct {
	repo *sql.DB
}

func NewSQLTypeRepository(repo *sql.DB) ports.ProductTypeRepository {
	return &mySQLTypeRepositoryAdapter{repo: repo}
}

func (r *mySQLTypeRepositoryAdapter) CreateType(newType domain.ProductType) ([]domain.ProductType, error) {
	query := "INSERT INTO product_types (name, active) VALUES (?, ?)"

	_, err := r.repo.Exec(query, newType.Name, newType.Active)
	if err != nil {
		return nil, err
	}

	ts, _ := r.GetAllTypes()
	return ts, nil
}

func (r *mySQLTypeRepositoryAdapter) GetType(id int) (*domain.ProductType, error) {
	query := fmt.Sprintf("SELECT * FROM product_types WHERE id = '%d'", id)
	result := r.repo.QueryRow(query)
	if result.Err() != nil {
		return nil, result.Err()
	}

	t := new(sqlProductType)
	if err := result.Scan(&t.id, &t.name, &t.active); err != nil {
		return nil, err
	}
	return t.ToDomain(), nil
}

func (r *mySQLTypeRepositoryAdapter) GetAllTypes() ([]domain.ProductType, error) {
	var ts []domain.ProductType
	query := "SELECT * FROM product_types"

	result, err := r.repo.Query(query)
	if err != nil {
		return nil, err
	}
	if result.Err() != nil {
		return nil, result.Err()
	}

	for result.Next() {
		t := sqlProductType{}

		if err := result.Scan(&t.id, &t.name, &t.active); err != nil {
			return nil, err
		}

		ts = append(ts, *t.ToDomain())
	}

	return ts, nil
}

func (r *mySQLTypeRepositoryAdapter) UpdateType(id int, update domain.ProductType) (*domain.ProductType, error) {
	t, _ := r.GetType(id)
	t.Update(update)

	query := "UPDATE product_types SET name = ?, active = ? WHERE id = ?"

	_, err := r.repo.Exec(query, t.Name, t.Active, id)
	if err != nil {
		return nil, err
	}

	t, _ = r.GetType(id)
	return t, nil
}

func (r *mySQLTypeRepositoryAdapter) DeleteType(id int, active bool) error {
	query := "UPDATE product_types SET active = ? WHERE id = ?"

	res, err := r.repo.Exec(query, active, id)
	if err != nil {
		return err
	}
	if count, err := res.RowsAffected(); !r.CheckIdExistence(id) {
		return errors.New(fmt.Sprintf(sqlErrorMessage+"could not delete, id %d not found", id))
	} else if count > 1 {
		log.Println(err)
		return errors.New(sqlErrorMessage + "multiple rows affected")
	}
	return nil
}

func (r *mySQLTypeRepositoryAdapter) CheckExistence(name string) bool {
	query := "SELECT * FROM product_types WHERE name = ?"

	res, err := r.repo.Exec(query, name)
	if err != nil {
		return false
	}
	if count, _ := res.RowsAffected(); count == 1 {
		return true
	}
	return false
}

func (r *mySQLTypeRepositoryAdapter) ValidateType(id int) bool {
	query := "SELECT * FROM product_types WHERE id = ? AND active = true"
	res := r.repo.QueryRow(query, id)
	t := new(sqlProductType)
	if err := res.Scan(&t.id, &t.name, &t.active); err != nil {
		return false
	}
	return true
}

func (r *mySQLTypeRepositoryAdapter) CheckIdExistence(id int) bool {
	query := "SELECT * FROM product_types WHERE id = ?"
	res := r.repo.QueryRow(query, id)
	if res.Err() != nil {
		return false
	}
	return true
}
