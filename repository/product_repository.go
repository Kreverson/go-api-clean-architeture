package repository

import (
	"database/sql"
	"fmt"
	"go-api/model"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}

func (p *ProductRepository) GetProducts() ([]model.Product, error) {

	query := "SELECT product_id, product_name, product_price FROM product"
	rows, err := p.connection.Query(query)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return []model.Product{}, err
	}

	var products []model.Product
	var product model.Product

	for rows.Next() {
		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Price)

		if err != nil {
			fmt.Println("Error scanning row:", err)
			return []model.Product{}, err
		}

		products = append(products, product)
	}

	rows.Close()

	return products, nil

}
