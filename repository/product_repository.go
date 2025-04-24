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

func (pr *ProductRepository) CreteProdutct(product model.Product) (model.Product, error) {

	query, err := pr.connection.Prepare("INSERT INTO product (product_name, product_price) VALUES ($1, $2) RETURNING product_id")

	if err != nil {
		fmt.Println("Error preparing query:", err)
		return model.Product{}, err
	}

	err = query.QueryRow(product.Name, product.Price).Scan(&product.ID)

	if err != nil {
		fmt.Println("Error executing query:", err)
		return model.Product{}, err
	}

	query.Close()

	return product, nil
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {

	query := "SELECT product_id, product_name, product_price FROM product"
	rows, err := pr.connection.Query(query)
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
