package service

import (
	"database/sql"
	"log"
	"transaction/repository"
)

type customerService struct {
	custRepo repository.CustomerRepository
}

func NewCustomerService(custRepo repository.CustomerRepository) CustomerService {
	return customerService{custRepo: custRepo}
}

func (s customerService) GetCustomers() ([]CustomerResponse, error) {

	customers, err := s.custRepo.GetAll()
	if err != nil {
		log.Println("Error while getting customers:", err)
		return nil, err
	}

	custResponses := []CustomerResponse{}
	for _, customer := range customers {
		custResponse := CustomerResponse{
			CustomerID: customer.CustomerID,
			Name:       customer.Name,
			Status:     customer.Status,
		}
		custResponses = append(custResponses, custResponse)
	}

	return custResponses, nil
}

func (s customerService) GetCustomer(id int) (*CustomerResponse, error) {
	customer, err := s.custRepo.GetById(id)
	if err != nil {

		if err == sql.ErrNoRows {
			return nil, err
		}

		log.Println("Error while getting customers:", err)
		return nil, err
	}

	custResponse := CustomerResponse{
		CustomerID: customer.CustomerID,
		Name:       customer.Name,
		Status:     customer.Status,
	}

	return &custResponse, nil
}

func (s customerService) PostCustomer(customerReq CustomerRequest) (*CustomerResponse, error) {
	customer := repository.Customer{
		Name:        customerReq.Name,
		DateOfBirth: customerReq.DateOfBirth,
		City:        customerReq.City,
		ZipCode:     customerReq.ZipCode,
		Status:      customerReq.Status,
	}

	_, err := s.custRepo.PostCustomer(customer)
	if err != nil {
		log.Println("Error while posting customer:", err)
		return nil, err
	}

	custResponse := CustomerResponse{
		Name:   customer.Name,
		Status: customer.Status,
	}

	return &custResponse, nil
}

func (s customerService) PutCustomer(customerReq CustomerRequest) (*CustomerResponse, error) {
	customer := repository.Customer{
		CustomerID:  customerReq.CustomerID,
		Name:        customerReq.Name,
		DateOfBirth: customerReq.DateOfBirth,
		City:        customerReq.City,
		ZipCode:     customerReq.ZipCode,
		Status:      customerReq.Status,
	}

	_, err := s.custRepo.PutCustomer(customer)
	if err != nil {
		log.Println("Error while putting customer:", err)
		return nil, err
	}

	custResponse := CustomerResponse{
		CustomerID: customer.CustomerID,
		Name:       customer.Name,
		Status:     customer.Status,
	}

	return &custResponse, nil
}

func (s customerService) DeleteCustomer(id int) (*CustomerResponse, error) {
	customer, err := s.custRepo.GetById(id)
	if err != nil {

		if err == sql.ErrNoRows {
			return nil, err
		}

		log.Println("Error while getting customers:", err)
		return nil, err
	}

	_, err = s.custRepo.DeleteCustomer(id)
	if err != nil {
		log.Println("Error while deleting customer:", err)
		return nil, err
	}

	custResponse := CustomerResponse{
		CustomerID: customer.CustomerID,
		Name:       customer.Name,
		Status:     customer.Status,
	}

	return &custResponse, nil
}
