package models

import (
	_ "github.com/go-sql-driver/mysql"
)

type Product struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

// Get all products from database
func GetAllProducts() []Product {
	db, err := connection()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	rows, err := db.Query("SELECT * FROM products")
	if err != nil {
		panic(err)
	}
	var products []Product
	for rows.Next() {
		var product Product
		err = rows.Scan(&product.ID, &product.Name, &product.Quantity, &product.Price)
		if err != nil {
			panic(err)
		}
		products = append(products, product)
	}
	return products
}

// Insert a new product into database
func InsertProducts(product Product) int64 {
	db, err := connection()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO products (name, quantity, price) VALUES (?, ?, ?)")
	if err != nil {
		panic(err)
	}
	res, err := stmt.Exec(product.Name, product.Quantity, product.Price)
	if err != nil {
		panic(err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}
	return id
}

// Get a product by id from database
func GetProduct(id int) Product {
	db, err := connection()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	row := db.QueryRow("SELECT * FROM products WHERE id = ?", id)
	var product Product
	err = row.Scan(&product.ID, &product.Name, &product.Quantity, &product.Price)
	if err != nil {
		panic(err)
	}

	return product
}
