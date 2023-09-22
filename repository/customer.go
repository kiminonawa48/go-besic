package repository

type Customer struct {
	CustomerID  int    `db:"customer_id"`
	Name        string `db:"name"`
	DateOfBirth string `db:"date_of_birth"`
	City        string `db:"city"`
	ZipCode     string `db:"zipcode"`
	Status      int    `db:"status"`
}

//go:generate mockgen -destination=../mock/mock_repository/mock_customer_repository.go bank/repository CustomerRepository
type CustomerRepository interface {
	GetAll() ([]Customer, error)
	GetById(int) (*Customer, error)
	PostCustomer(Customer) (int64, error)
	PutCustomer(Customer) (int64, error)
	DeleteCustomer(int) (int64, error)
}
