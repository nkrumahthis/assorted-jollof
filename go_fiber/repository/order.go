package repository

import (
	"nkrumahthis/assorted-jollof/database"
	"time"
)

type Order struct {
	ID         int    `json:"id"`
	Packs      int    `json:"packs"`
	CustomerId int    `json:"customer_id"`
	Location   string `json:"location"`
	Status     string `json:"status"`
	CreatedAt  any    `json:"created_at"`
}

func CreateOrder(packs int, customer_id int, location string, status string) (*Order, error) {
	// fetch database
	db, err := database.GetDB()
	if err != nil {
		return nil, err
	}

	// prepare insert query
	query := "INSERT INTO orders (id, packs, customer_id, location, status, created_at) VALUES ($1, $2, $3, $4, $5, $6)"
	result, err := db.Exec(query, nil, packs, customer_id, location, status, time.Now())
	if err != nil {
		return nil, err
	}

	// get last insert id
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	order, err := FindOrder(int(id))
	if err != nil {
		return nil, err
	}

	return &order, err
}

func FindAllOrders() ([]Order, error) {
	db, err := database.GetDB()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT * from orders")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orders := make([]Order, 0)
	for rows.Next() {
		var order Order
		err := rows.Scan(&order.ID, &order.Packs, &order.CustomerId, &order.Location, &order.Status, &order.CreatedAt)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func FindOrdersByParams(id, packs, customerId, location, status, createdAt string) ([]Order, error) {
	// get an instance of the database
	db, err := database.GetDB()
	if err != nil {
		return nil, err
	}

	query := "SELECT * FROM orders WHERE 1 = 1" // Start with a true condition
	var params []interface{}                    // slice to hold parameter values

	// Add the filter conditions and parameters based on the provided parameters
	if id != "" {
		query += " AND id = ?"
		params = append(params, id)
	}
	if packs != "" {
		query += " AND packs = ?"
		params = append(params, packs)
	}
	if customerId != "" {
		query += " AND customer_id = ?"
		params = append(params, customerId)
	}
	if location != "" {
		query += " AND location = ?"
		params = append(params, location)
	}
	if status != "" {
		query += " AND status = ?"
		params = append(params, status)
	}
	if createdAt != "" {
		query += " AND created_at = ?"
		params = append(params, createdAt)
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

	// Iterate over the rows and scan the data into Order structs
	var orders []Order
	for rows.Next() {
		var order Order
		err = rows.Scan(&order.ID, &order.Packs, &order.CustomerId, &order.Location, &order.Status, &order.CreatedAt)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil

}

func FindOrder(id int) (Order, error) {
	var order Order

	db, err := database.GetDB()
	if err != nil {
		return order, err
	}

	row := db.QueryRow("SELECT * from orders WHERE id=$1", id)
	err = row.Scan(&order.ID, &order.Packs, &order.CustomerId, &order.Location, &order.Status, &order.CreatedAt)

	return order, err

}
