package repository

import (
	"database/sql"
	"go-kafka-arquitetura-hexagonal /internal/domain/entity"
)

type ProductRepositoryMsql struct {
	DB *sql.DB
}

func NewProductRepositoryMysql(db *sql.DB) *ProductRepositoryMsql {
	return &ProductRepositoryMsql{
		DB: db,
	}
}

func (r *ProductRepositoryMsql) Create(product *entity.Product) error {
	_, err := r.DB.Exec("Insert into products (id, name, price) values(?,?,?)", product.ID, product.Name, product.Price)

	if err != nil {
		return err
	}

	return nil
}

func (r *ProductRepositoryMsql) FindAll() ([]*entity.Product, error) {
	rows, err := r.DB.Query("Select * from products")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []*entity.Product

	for rows.Next() {
		var product entity.Product
		err = rows.Scan(&product.ID, &product.Name, &product.Price)

		if err != nil {
			return nil, err
		}

		products = append(products, &product)
	}

	return products, nil
}
