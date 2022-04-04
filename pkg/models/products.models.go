package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/pauloeduardods/e-commerce/pkg/schemas"
)

// Get all products from database
func GetAllProducts(products chan []schemas.Product) {
	db, err := connection()
	if err != nil {
		products <- []schemas.Product{}
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM products")
	if err != nil {
		products <- []schemas.Product{}
		return
	}
	var productsSplice []schemas.Product
	for rows.Next() {
		var product schemas.Product
		err = rows.Scan(&product.ID, &product.Name, &product.Quantity, &product.Price)
		if err != nil {
			products <- []schemas.Product{}
			return
		}
		productsSplice = append(productsSplice, product)
	}
	if productsSplice == nil {
		products <- []schemas.Product{}
		return
	}
	products <- productsSplice
}

// Insert a new product into database
func InsertProduct(product schemas.Product, id chan int64) {
	db, err := connection()
	if err != nil {
		id <- 0
		return
	}

	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO products (name, quantity, price) VALUES (?, ?, ?)")
	if err != nil {
		id <- 0
		return
	}
	res, err := stmt.Exec(product.Name, product.Quantity, product.Price)
	if err != nil {
		id <- 0
		return
	}
	insertedId, err := res.LastInsertId()
	if err != nil {
		id <- 0
		return
	}
	id <- insertedId
}

// Get a product by id from database
func GetProduct(id int, product chan schemas.Product) {
	db, err := connection()
	if err != nil {
		product <- schemas.Product{}
		return
	}

	defer db.Close()

	row := db.QueryRow("SELECT * FROM products WHERE id = ?", id)
	var curProduct schemas.Product
	err = row.Scan(&curProduct.ID, &curProduct.Name, &curProduct.Quantity, &curProduct.Price)
	if curProduct == (schemas.Product{}) || err != nil {
		product <- schemas.Product{}
		return
	}
	product <- curProduct
}
