package service

type TransactionResponse struct {
	TransactionID int    `json:"transaction_id"`
	BankID        int    `json:"bank_id"`
	CustomerID    int    `json:"customer_id"`
	Amount        int    `json:"amount"`
	Status        int    `json:"status"`
	BankName      string `json:"bank_name"`
	CustomerName  string `json:"customer_name"`
}

type TransactionRequest struct {
	TransactionID int    `json:"transaction_id"`
	BankID        int    `json:"bank_id"`
	CustomerID    int    `json:"customer_id"`
	Amount        int    `json:"amount"`
	Status        int    `json:"status"`
	BankName      string `json:"bank_name"`
	CustomerName  string `json:"customer_name"`
}

type TransactionService interface {
	GetTransactionsAll() ([]TransactionResponse, error)
	GetTransactionByCustomerId(int) ([]TransactionResponse, error)
	GetTransactionByBankId(int) ([]TransactionResponse, error)
	PostTransaction(TransactionRequest) (*TransactionResponse, error)
	PutTransaction(TransactionRequest) (*TransactionResponse, error)
	DeleteTransaction(int) (*TransactionResponse, error)
}
