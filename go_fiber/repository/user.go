package repository

import (
	"nkrumahthis/assorted-jollof/database"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func FindAllUsers() ([]User, error) {
	db,err := database.GetDB()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query(("SELECT * from users"))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func FindUser(id int) (User, error) {
	var user User

	db, err := database.GetDB()
	if err != nil {
		return user, err 
	}

	row := db.QueryRow("SELECT * FROM users WHERE id = ?", id)
	err = row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)

	return user, err
}
