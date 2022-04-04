package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/pauloeduardods/e-commerce/pkg/schemas"
)

// Get all products from database
func GetAllProducts() ([]schemas.Product, error) {
	db, err := connection()
	if err != nil {
		return nil, err
	}

	defer db.Close()

	rows, err := db.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}
	var products []schemas.Product
	for rows.Next() {
		var product schemas.Product
		err = rows.Scan(&product.ID, &product.Name, &product.Quantity, &product.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, err
}

// Insert a new product into database
func InsertProducts(product schemas.Product) (int64, error) {
	db, err := connection()
	if err != nil {
		return 0, err
	}

	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO products (name, quantity, price) VALUES (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(product.Name, product.Quantity, product.Price)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

// Get a product by id from database
func GetProduct(id int) (schemas.Product, error) {
	db, err := connection()
	if err != nil {
		return schemas.Product{}, err
	}

	defer db.Close()

	row := db.QueryRow("SELECT * FROM products WHERE id = ?", id)
	var product schemas.Product
	err = row.Scan(&product.ID, &product.Name, &product.Quantity, &product.Price)
	if err != nil {
		return schemas.Product{}, err
	}

	return product, nil
}
