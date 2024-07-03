package repositories

import (
	"api-produtos/internal/core/domain"
	"api-produtos/internal/core/ports"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

const sqlErrorMessage = "MySQLRepo error: "

type mySQLProductRepositoryAdapter struct {
	repo *sql.DB
}

func NewSQLProductRepository(db *sql.DB) ports.ProductRepository {
	return &mySQLProductRepositoryAdapter{repo: db}
}

func (r *mySQLProductRepositoryAdapter) CreateProduct(product domain.Product) error {
	query := "INSERT INTO products (name, description, price, type_id) VALUES (?, ?, ?, ?)"

	_, err := r.repo.Exec(query, product.Name, product.Description, product.Price, product.Type.Id)
	if err != nil {
		return err
	}

	return nil
}

func (r *mySQLProductRepositoryAdapter) GetProduct(id int) (*domain.Product, error) {
	query := fmt.Sprintf("SELECT * FROM products WHERE id = '%d'", id)
	result := r.repo.QueryRow(query)
	if result.Err() != nil {
		return nil, result.Err()
	}

	p := new(sqlProduct)
	if err := result.Scan(&p.id, &p.name, &p.description, &p.price, &p.typeId); err != nil {
		return nil, err
	}
	return p.ToDomain(), nil
}

func (r *mySQLProductRepositoryAdapter) GetAllProducts() ([]domain.Product, error) {
	var ps []domain.Product
	query := "SELECT * FROM products"

	result, err := r.repo.Query(query)
	if err != nil {
		return nil, err
	}
	if result.Err() != nil {
		return nil, result.Err()
	}

	for result.Next() {
		p := sqlProduct{}

		if err := result.Scan(&p.id, &p.name, &p.description, &p.price, &p.typeId); err != nil {
			return nil, err
		}

		ps = append(ps, *p.ToDomain())
	}

	return ps, nil
}

func (r *mySQLProductRepositoryAdapter) UpdateProduct(id int, update domain.Product) (*domain.Product, error) {
	p, _ := r.GetProduct(id)
	p.Update(update)

	query := "UPDATE products SET name = ?, description = ?, price = ? WHERE id = ?"

	_, err := r.repo.Exec(query, p.Name, p.Description, p.Price, id)
	if err != nil {
		return nil, err
	}

	p, _ = r.GetProduct(id)
	return p, nil
}

func (r *mySQLProductRepositoryAdapter) DeleteProduct(id int) error {
	query := "DELETE FROM products WHERE id = ?"

	res, err := r.repo.Exec(query, id)
	if err != nil {
		return err
	}
	if count, err := res.RowsAffected(); count == 0 {
		return errors.New(fmt.Sprintf(sqlErrorMessage+"could not delete, id %d not found", id))
	} else if count != 1 {
		log.Println(err)
		return errors.New(sqlErrorMessage + "multiple rows affected")
	}
	return nil
}

func buildMysqlConnUrl() string {
	user := os.Getenv("MYSQL_USR")
	pass := os.Getenv("MYSQL_ROOT_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	dbName := os.Getenv("MYSQL_DATABASE")

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=true", user, pass, host, port, dbName)
}

func StartMysqlDb() *sql.DB {
	db, err := sql.Open("mysql", buildMysqlConnUrl())
	if err != nil {
		panic(err)
	}

	return db
}
