package service

import (
	"database/sql"
	"log"
	"transaction/repository"
)

type transactionService struct {
	transactionRepo repository.TransactionRepository
}

func NewTransactionService(transactionRepo repository.TransactionRepository) TransactionService {
	return transactionService{transactionRepo: transactionRepo}
}

func (s transactionService) GetTransactionsAll() ([]TransactionResponse, error) {

	transactions, err := s.transactionRepo.GetTransactionsAll()
	if err != nil {
		log.Println("Error while getting transactions:", err)
		return nil, err
	}

	transactionResponses := []TransactionResponse{}
	for _, transaction := range transactions {
		transactionResponse := TransactionResponse{
			TransactionID: transaction.TransactionID,
			BankID:        transaction.BankID,
			CustomerID:    transaction.CustomerID,
			Amount:        transaction.Amount,
			Status:        transaction.Status,
			BankName:      transaction.BankName,
			CustomerName:  transaction.CustomerName,
		}
		transactionResponses = append(transactionResponses, transactionResponse)
	}

	return transactionResponses, nil
}

func (s transactionService) GetTransactionByCustomerId(id int) ([]TransactionResponse, error) {
	transactions, err := s.transactionRepo.GetTransactionByCustomerId(id)
	if err != nil {

		if err == sql.ErrNoRows {
			return nil, err
		}

		log.Println("Error while getting transactions:", err)
		return nil, err
	}

	transactionResponses := []TransactionResponse{}
	for _, transaction := range transactions {
		transactionResponse := TransactionResponse{
			TransactionID: transaction.TransactionID,
			BankID:        transaction.BankID,
			CustomerID:    transaction.CustomerID,
			Amount:        transaction.Amount,
			Status:        transaction.Status,
			BankName:      transaction.BankName,
			CustomerName:  transaction.CustomerName,
		}
		transactionResponses = append(transactionResponses, transactionResponse)
	}

	return transactionResponses, nil
}

func (s transactionService) GetTransactionByBankId(id int) ([]TransactionResponse, error) {
	transactions, err := s.transactionRepo.GetTransactionByBankId(id)
	if err != nil {

		if err == sql.ErrNoRows {
			return nil, err
		}

		log.Println("Error while getting transactions:", err)
		return nil, err
	}

	transactionResponses := []TransactionResponse{}
	for _, transaction := range transactions {
		transactionResponse := TransactionResponse{
			TransactionID: transaction.TransactionID,
			BankID:        transaction.BankID,
			CustomerID:    transaction.CustomerID,
			Amount:        transaction.Amount,
			Status:        transaction.Status,
			BankName:      transaction.BankName,
			CustomerName:  transaction.CustomerName,
		}
		transactionResponses = append(transactionResponses, transactionResponse)
	}

	return transactionResponses, nil
}

func (s transactionService) PostTransaction(transactionReq TransactionRequest) (*TransactionResponse, error) {

	transaction := repository.TransactionModel{
		BankID:     transactionReq.BankID,
		CustomerID: transactionReq.CustomerID,
		Amount:     transactionReq.Amount,
		Status:     transactionReq.Status,
	}

	newTransaction, err := s.transactionRepo.PostTransaction(transaction)

	if err != nil {
		log.Println("Error while posting transaction:", err)
		return nil, err
	}

	transactionResponse := TransactionResponse{
		TransactionID: newTransaction.TransactionID,
		BankID:        newTransaction.BankID,
		CustomerID:    newTransaction.CustomerID,
		Amount:        newTransaction.Amount,
		Status:        newTransaction.Status,
	}

	return &transactionResponse, nil
}

func (s transactionService) PutTransaction(transactionReq TransactionRequest) (*TransactionResponse, error) {

	transaction := repository.TransactionModel{
		TransactionID: transactionReq.TransactionID,
		BankID:        transactionReq.BankID,
		CustomerID:    transactionReq.CustomerID,
		Amount:        transactionReq.Amount,
		Status:        transactionReq.Status,
	}

	_, err := s.transactionRepo.PutTransaction(transaction)
	if err != nil {
		log.Println("Error while putting transaction:", err)
		return nil, err
	}

	transactionResponse := TransactionResponse{
		TransactionID: transaction.TransactionID,
		BankID:        transaction.BankID,
		CustomerID:    transaction.CustomerID,
		Amount:        transaction.Amount,
		Status:        transaction.Status,
	}

	return &transactionResponse, nil
}

func (s transactionService) DeleteTransaction(id int) (*TransactionResponse, error) {
	_, err := s.transactionRepo.DeleteTransaction(id)

	if err != nil {
		log.Println("Error while deleting transaction:", err)
		return nil, err
	}

	TransactionResponse := TransactionResponse{
		TransactionID: id,
	}

	return &TransactionResponse, nil

}
