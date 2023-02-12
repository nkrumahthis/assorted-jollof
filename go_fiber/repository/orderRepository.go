package model

import "nkrumahthis/assorted-jollof/database"

type Order struct {
	ID         int    `json:"id"`
	Packs      int    `json:"packs"`
	CustomerId int    `json:"customer_id"`
	Location   string `json:"location"`
	Status     string `json:"status"`
	CreatedAt  any `json:"created_at"`
}

func FindAll() ([]Order, error) {
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
		err := rows.Scan(&order.ID, &order.Packs, &order.CustomerId, &order.Location, &order.Status, &order.CreatedAt); 
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

func find(id int) (Order, error) {
	var order Order

	db, err := database.GetDB()
	if err != nil {
		return order, err
	}

	row := db.QueryRow("SELECT * from order WHERE id=$1", id)
	err = row.Scan(&order.ID, &order.Packs, &order.CustomerId, &order.Location, &order.Status, &order.CreatedAt); 

	return order, err

}
