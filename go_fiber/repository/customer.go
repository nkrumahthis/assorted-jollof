package repository

import "nkrumahthis/assorted-jollof/database"

type Customer struct{
	ID int 
	Name string 
	Phone string
	Token string
}

func FindAllCustomers() ([]Customer, error){
	db, err := database.GetDB()
	if err != nil{
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

func FindCustomer(id int) (Customer, error){
	var customer Customer

	db, err := database.GetDB()
	if err != nil {
		return customer, err
	}

	row := db.QueryRow("SELECT * from customers WHERE id=$1", id)
	err = row.Scan(&customer.ID, &customer.Name, &customer.Phone, &customer.Token)

	return customer, err
}