package models

import "github.com/pauloeduardods/e-commerce/pkg/schemas"

func GetUserByEmail(email string) (schemas.User, error) {
	db, err := connection()
	if err != nil {
		return schemas.User{}, err
	}

	defer db.Close()

	row := db.QueryRow("SELECT * FROM users WHERE email = ?", email)
	var user schemas.User
	err = row.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return schemas.User{}, err
	}

	return user, nil
}

func CreateUser(user schemas.User) (int64, error) {
	db, err := connection()
	if err != nil {
		return 0, err
	}

	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO users (username, email, password) VALUES (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(user.Username, user.Email, user.Password)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}
