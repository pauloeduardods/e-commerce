package models

import "github.com/pauloeduardods/e-commerce/pkg/schemas"

func GetUserByEmail(email string, user chan schemas.User) {
	db, err := connection()
	if err != nil {
		return
	}

	defer db.Close()

	row := db.QueryRow("SELECT * FROM users WHERE email = ?", email)
	var curUser schemas.User
	err = row.Scan(&curUser.ID, &curUser.Username, &curUser.Email, &curUser.Password)
	if err != nil {
		return
	}
	user <- curUser
}

func CreateUser(user schemas.User, id chan int64) {
	db, err := connection()
	if err != nil {
		return
	}

	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO users (username, email, password) VALUES (?, ?, ?)")
	if err != nil {
		return
	}
	res, err := stmt.Exec(user.Username, user.Email, user.Password)
	if err != nil {
		return
	}
	insertedId, err := res.LastInsertId()
	if err != nil {
		return
	}
	id <- insertedId
}
