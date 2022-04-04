package models

import "github.com/pauloeduardods/e-commerce/pkg/schemas"

func GetUserByEmail(email string, user chan schemas.User) {
	db, err := connection()
	if err != nil {
		user <- schemas.User{}
		return
	}

	defer db.Close()

	row := db.QueryRow("SELECT * FROM users WHERE email = ?", email)
	var curUser schemas.User
	err = row.Scan(&curUser.ID, &curUser.Username, &curUser.Email, &curUser.Password)
	if curUser == (schemas.User{}) || err != nil {
		user <- schemas.User{}
		return
	}
	user <- curUser
}

func CreateUser(user schemas.User, id chan int64) {
	db, err := connection()
	if err != nil {
		id <- 0
		return
	}

	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO users (username, email, password) VALUES (?, ?, ?)")
	if err != nil {
		id <- 0
		return
	}
	res, err := stmt.Exec(user.Username, user.Email, user.Password)
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
