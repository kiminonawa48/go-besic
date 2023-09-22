package repository

import "github.com/jmoiron/sqlx"

type transactionRepositoryDB struct {
	db *sqlx.DB
}

func NewTransactionRepositoryDB(db *sqlx.DB) transactionRepositoryDB {
	return transactionRepositoryDB{db: db}
}

func (d transactionRepositoryDB) GetTransactionsAll() ([]TransactionModel, error) {
	transactions := []TransactionModel{}
	// join with customer and bank
	findAllSql := "select t.transaction_id, t.bank_id, t.customer_id, t.amount, t.status, b.name as bank_name, c.name as customer_name from transactions t join banks b on t.bank_id = b.bank_id join customers c on t.customer_id = c.customer_id"

	err := d.db.Select(&transactions, findAllSql)

	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (d transactionRepositoryDB) GetTransactionByCustomerId(id int) ([]TransactionModel, error) {
	transactions := []TransactionModel{}
	// join with customer and bank
	findAllSql := "select t.transaction_id, t.bank_id, t.customer_id, t.amount, t.status, b.name as bank_name, c.name as customer_name from transactions t join banks b on t.bank_id = b.bank_id join customers c on t.customer_id = c.customer_id where t.customer_id = ?"

	err := d.db.Select(&transactions, findAllSql, id)

	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (d transactionRepositoryDB) GetTransactionByBankId(id int) ([]TransactionModel, error) {
	transactions := []TransactionModel{}
	// join with customer and bank
	findAllSql := "select t.transaction_id, t.bank_id, t.customer_id, t.amount, t.status, b.name as bank_name, c.name as customer_name from transactions t join banks b on t.bank_id = b.bank_id join customers c on t.customer_id = c.customer_id where t.bank_id = ?"

	err := d.db.Select(&transactions, findAllSql, id)

	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (d transactionRepositoryDB) PostTransaction(transaction TransactionModel) (*TransactionModel, error) {
	insertSql := "insert into transactions (bank_id, customer_id, amount, status) values (?, ?, ?, ?)"

	result, err := d.db.Exec(insertSql, transaction.BankID, transaction.CustomerID, transaction.Amount, transaction.Status)

	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return nil, err
	}

	transaction.TransactionID = int(id)

	return &transaction, nil
}
func (d transactionRepositoryDB) PutTransaction(transaction TransactionModel) (int64, error) {
	updateSql := "update transactions set bank_id = ?, customer_id = ?, amount = ?, status = ? where transaction_id = ?"

	result, err := d.db.Exec(updateSql, transaction.BankID, transaction.CustomerID, transaction.Amount, transaction.Status, transaction.TransactionID)

	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

func (d transactionRepositoryDB) DeleteTransaction(id int) (int64, error) {
	deleteSql := "delete from transactions where transaction_id = ?"

	result, err := d.db.Exec(deleteSql, id)

	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}
