package models

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func connection() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:paulopaulo@tcp(localhost:3306)/Go_Ecommerce")
	if err != nil {
		return db, err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db, nil
}

// func execSomeQuery() {
// 	res := conn().Exec("SELECT * FROM users")
// 	fmt.Println(res)
// }
