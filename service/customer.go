package service

type CustomerResponse struct {
	CustomerID int    `json:"customer_id"`
	Name       string `json:"name"`
	Status     int    `json:"status"`
}

type CustomerRequest struct {
	CustomerID  int    `json:"customer_id"`
	Name        string `json:"name"`
	DateOfBirth string `json:"date_of_birth"`
	City        string `json:"city"`
	ZipCode     string `json:"zipcode"`
	Status      int    `json:"status"`
}

type CustomerService interface {
	GetCustomers() ([]CustomerResponse, error)
	GetCustomer(int) (*CustomerResponse, error)
	PostCustomer(CustomerRequest) (*CustomerResponse, error)
	PutCustomer(CustomerRequest) (*CustomerResponse, error)
	DeleteCustomer(int) (*CustomerResponse, error)
}
