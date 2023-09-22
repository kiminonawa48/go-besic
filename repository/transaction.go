package repository

// transaction with join table bank and customer
type TransactionModel struct {
	TransactionID int    `db:"transaction_id"`
	BankID        int    `db:"bank_id"`
	CustomerID    int    `db:"customer_id"`
	Amount        int    `db:"amount"`
	Status        int    `db:"status"`
	BankName      string `db:"bank_name"`
	CustomerName  string `db:"customer_name"`
}

type TransactionRepository interface {
	GetTransactionsAll() ([]TransactionModel, error)
	GetTransactionByCustomerId(int) ([]TransactionModel, error)
	GetTransactionByBankId(int) ([]TransactionModel, error)
	PostTransaction(TransactionModel) (*TransactionModel, error)
	PutTransaction(TransactionModel) (int64, error)
	DeleteTransaction(int) (int64, error)
}
