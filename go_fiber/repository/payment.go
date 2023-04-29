package repository

import (
	"nkrumahthis/assorted-jollof/database"
	"time"
)

type Payment struct {
	ID         int        `json:"id"`
	Amount     float32    `json:"amount"`
	CustomerId int        `json:"customer_id"`
	OrderId		int 	`json:"order_id"`
	CreatedAt  *time.Time `json:"created_at"`
	UserId     int        `json:"user_id"`
}

func CreatePayment(orderId int, amount float32, customerId int) (*Payment, error) {
	db, err := database.GetDB()

	if err != nil {
		return nil, err
	}

	query := "INSERT INTO payments (id, order_id, amount, customer_id, created_at) VALUES ($1, $2, $3, $4, $5)"
	result, err := db.Exec(query, nil, orderId, amount, customerId, time.Now())
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	payment, err := FindPayment(int(id))
	if err != nil {
		return nil, err
	}

	return &payment, err
}

func FindPayment(id int) (Payment, error) {
	var payment Payment
	db, err := database.GetDB()
	if err != nil {
		return payment, err
	}

	row := db.QueryRow("SELECT * FROM payments WHERE id=$1", id)
	err = row.Scan(&payment.ID, &payment.CustomerId, &payment.Amount, &payment.UserId, &payment.CreatedAt)
	return payment, err
}

func FindPaymentsByParams(id, amount, orderId, timestamp, userId string) ([]Payment, error) {
	db, err := database.GetDB()
	if err != nil {
		return nil, err
	}

	query := "SELECT * from payments WHERE 1 = 1" // start with a true condition
	var params []interface{}

	// add the filter conditions and parameters based on the payment parameters
	if id != "" {
		query += " AND id = ?"
		params = append(params, id)
	}
	if amount != "" {
		query += " AND amount = ?"
		params = append(params, amount)
	}
	if orderId != "" {
		query += " AND order_id = ?"
		params = append(params, orderId)
	}
	if timestamp != "" {
		query += " AND timestamp = ?"
		params = append(params, timestamp)
	}
	if userId != "" {
		query += " AND user_id = ?"
		params = append(params, userId)
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

	var payments []Payment
	for rows.Next() {
		var payment Payment
		err = rows.Scan(&payment.ID, &payment.Amount, &payment.CustomerId, &payment.UserId, &payment.CreatedAt)
		if err != nil {
			return nil, err
		}
		payments = append(payments, payment)
	}

	return payments, nil
}
