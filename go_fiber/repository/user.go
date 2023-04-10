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

func FindUsersWithFilter(name, email, password string) ([]User, error) {
	var users []User

	db, err := database.GetDB()
	if err != nil {
		return users,err
	}

	// Build the SQL query based on the provided filter parameters
    query := "SELECT * FROM users WHERE 1 = 1" // Start with a true condition

    // Add the filter conditions based on the provided parameters
    if email != "" {
        query += " AND email = ?"
    }
    if name != "" {
        query += " AND name = ?"
    }
    if password != "" {
        query += " AND password = ?"
    }

    // Prepare the SQL statement
    stmt, err := db.Prepare(query)
    if err != nil {
        // Handle error
        return users, err
    }
    defer stmt.Close()

    // Execute the SQL statement with the provided filter parameters
    rows, err := stmt.Query(email, name, password)
    if err != nil {
        // Handle error
        return users, err
    }
    defer rows.Close()

    // Iterate over the rows and scan the data into User structs
    for rows.Next() {
        var user User
        err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
        if err != nil {
            // Handle error
            return users, err
        }
        users = append(users, user)
    }

    // Return the users as JSON
    // You can use the c.JSON() function to send the users as JSON response
    return users, nil
}
