package repository

import (
	"nkrumahthis/assorted-jollof/database"
)

// User struct.
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Find all Users in the database.
func FindAllUsers() ([]User, error) {
	db, err := database.GetDB()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query(("SELECT * from users where 1=1"))
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

// Find User by id.
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

func FindUsersWithParams(name, email, password string) ([]User, error) {

	db, err := database.GetDB()
	if err != nil {
		return nil,err
	}

	query := "SELECT * FROM users WHERE 1 = 1" // Start with a true condition
	var params []interface{}                // Slice to hold the parameter values

	// Add the filter conditions and parameters based on the provided parameters
	if email != "" {
		query += " AND email = ?"
		params = append(params, email)
	}
	if name != "" {
		query += " AND name = ?"
		params = append(params, name)
	}
	if password != "" {
		query += " AND password = ?"
		params = append(params, password)
	}

	// Prepare the SQL statement
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// Execute the SQL statement with the provided filter parameters
	rows, err := stmt.Query(params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate over the rows and scan the data into User structs
	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
