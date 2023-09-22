package repository

import "github.com/jmoiron/sqlx"

type customerRepositoryDB struct {
	db *sqlx.DB
}

func NewCustomerRepositoryDB(db *sqlx.DB) customerRepositoryDB {
	return customerRepositoryDB{db: db}
}

func (d customerRepositoryDB) GetAll() ([]Customer, error) {
	customers := []Customer{}
	findAllSql := "select customer_id, name, date_of_birth, city, zipcode, status from customers"

	err := d.db.Select(&customers, findAllSql)

	if err != nil {
		return nil, err
	}

	return customers, nil
}

func (d customerRepositoryDB) GetById(id int) (*Customer, error) {
	customer := Customer{}
	findByIdSql := "select customer_id, name, date_of_birth, city, zipcode, status from customers where customer_id = ?"

	err := d.db.Get(&customer, findByIdSql, id)

	if err != nil {
		return nil, err
	}

	return &customer, nil
}

func (d customerRepositoryDB) PostCustomer(customer Customer) (int64, error) {
	insertSql := "insert into customers (name, date_of_birth, city, zipcode, status) values (?, ?, ?, ?, ?)"

	result, err := d.db.Exec(insertSql, customer.Name, customer.DateOfBirth, customer.City, customer.ZipCode, customer.Status)

	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

func (d customerRepositoryDB) PutCustomer(customer Customer) (int64, error) {
	updateSql := "update customers set name = ?, date_of_birth = ?, city = ?, zipcode = ?, status = ? where customer_id = ?"

	result, err := d.db.Exec(updateSql, customer.Name, customer.DateOfBirth, customer.City, customer.ZipCode, customer.Status, customer.CustomerID)

	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

func (d customerRepositoryDB) DeleteCustomer(id int) (int64, error) {
	deleteSql := "delete from customers where customer_id = ?"

	result, err := d.db.Exec(deleteSql, id)

	if err != nil {
		return 0, err
	}

	return result.RowsAffected()

}
