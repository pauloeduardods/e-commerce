package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/pauloeduardods/e-commerce/pkg/schemas"
)

// Get all products from database
func GetAllProducts(products chan []schemas.Product) {
	db, err := connection()
	if err != nil {
		return
	}

	defer db.Close()

	rows, err := db.Query("SELECT * FROM products")
	if err != nil {
		return
	}
	var productsSplice []schemas.Product
	for rows.Next() {
		var product schemas.Product
		err = rows.Scan(&product.ID, &product.Name, &product.Quantity, &product.Price)
		if err != nil {
			return
		}
		productsSplice = append(productsSplice, product)
	}
	products <- productsSplice
}

// Insert a new product into database
func InsertProducts(product schemas.Product, id chan int64) {
	db, err := connection()
	if err != nil {
		id <- 0
	}

	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO products (name, quantity, price) VALUES (?, ?, ?)")
	if err != nil {
		id <- 0
	}
	res, err := stmt.Exec(product.Name, product.Quantity, product.Price)
	if err != nil {
		id <- 0
	}
	insertedId, err := res.LastInsertId()
	if err != nil {
		id <- 0
	}
	id <- insertedId
}

// Get a product by id from database
func GetProduct(id int, product chan schemas.Product) {
	db, err := connection()
	if err != nil {
		return
	}

	defer db.Close()

	row := db.QueryRow("SELECT * FROM products WHERE id = ?", id)
	var curProduct schemas.Product

	err = row.Scan(&curProduct.ID, &curProduct.Name, &curProduct.Quantity, &curProduct.Price)
	if err != nil {
		return
	}
	product <- curProduct
}
