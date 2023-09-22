package repository

type BankModel struct {
	BankID int    `db:"bank_id"`
	Name   string `db:"name"`
	Status int    `db:"status"`
}

type BankRepository interface {
	GetBanksAll() ([]BankModel, error)
	GetBankById(int) (*BankModel, error)
	PostBank(BankModel) (*BankModel, error)
	PutBank(BankModel) (int64, error)
	DeleteBank(int) (int64, error)
}
