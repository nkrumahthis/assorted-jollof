package repository

import "nkrumahthis/assorted-jollof/database"

type Customer struct {
	ID    int
	Name  string
	Phone string
	Token string
}

func CreateCustomer(name string, phone string, token string) (*Customer, error) {
	// get database
	db, err := database.GetDB()
	if err != nil {
		return nil, err
	}

	// create query, execute
	query := "INSERT INTO customers (id, name, phone, token) VALUES ($1, $2, $3, $4)"
	result, err := db.Exec(query, nil, name, phone, token)
	if err != nil {
		return nil, err
	}

	// get id of new row
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	// retrieve new customer
	customer, err := FindCustomer(int(id))
	return &customer, err
}

func FindAllCustomers() ([]Customer, error) {
	db, err := database.GetDB()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT * from customers")
	if err != nil {
		return nil, err
	}

	customers := make([]Customer, 0)
	for rows.Next() {
		var customer Customer
		err := rows.Scan(&customer.ID, &customer.Name, &customer.Phone, &customer.Token)
		if err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}

	return customers, nil

}

func FindCustomer(id int) (Customer, error) {
	var customer Customer

	db, err := database.GetDB()
	if err != nil {
		return customer, err
	}

	row := db.QueryRow("SELECT * from customers WHERE id=$1", id)
	err = row.Scan(&customer.ID, &customer.Name, &customer.Phone, &customer.Token)

	return customer, err
}

func FindCustomersByParams(id, name, phone, token string) ([]Customer, error) {
	// get an instance of the database
	db, err := database.GetDB()
	if err != nil {
		return nil, err
	}

	query := "SELECT * FROM customers WHERE 1 = 1" // Start with a true condition
	var params []interface{}

	// Add the filter conditions and parameters based on the providers parameters
	if id != "" {
		query += " AND id = ?"
		params = append(params, id)
	}
	if name != "" {
		query += " AND names = ?"
		params = append(params, name)
	}
	if phone != "" {
		query += " AND phone = ?"
		params = append(params, phone)
	}
	if token != "" {
		query += " AND token = ?"
		params = append(params, token)
	}

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customers []Customer
	for rows.Next() {
		var customer Customer
		err = rows.Scan(&customer.ID, &customer.Name, &customer.Phone, &customer.Token)
		if err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}

	return customers, nil
}
