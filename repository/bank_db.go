package repository

import "github.com/jmoiron/sqlx"

type bankRepositoryDB struct {
	db *sqlx.DB
}

func NewBankRepositoryDB(db *sqlx.DB) bankRepositoryDB {
	return bankRepositoryDB{db: db}
}

func (d bankRepositoryDB) GetBanksAll() ([]BankModel, error) {
	banks := []BankModel{}
	findAllSql := "select bank_id, name, status from banks"

	err := d.db.Select(&banks, findAllSql)

	if err != nil {
		return nil, err
	}

	return banks, nil
}
func (d bankRepositoryDB) GetBankById(id int) (*BankModel, error) {
	bank := BankModel{}

	findByIdSql := "select bank_id, name, status from banks where bank_id = ?"

	err := d.db.Get(&bank, findByIdSql, id)

	if err != nil {
		return nil, err
	}

	return &bank, nil
}

func (d bankRepositoryDB) PostBank(bank BankModel) (*BankModel, error) {
	insertSql := "insert into banks (name, status) values (?, ?)"

	result, err := d.db.Exec(insertSql, bank.Name, bank.Status)

	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return nil, err
	}

	bank.BankID = int(id)

	return &bank, nil
}

func (d bankRepositoryDB) PutBank(bank BankModel) (int64, error) {
	updateSql := "update banks set name = ?, status = ? where bank_id = ?"

	result, err := d.db.Exec(updateSql, bank.Name, bank.Status, bank.BankID)

	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

func (d bankRepositoryDB) DeleteBank(id int) (int64, error) {
	deleteSql := "delete from banks where bank_id = ?"

	result, err := d.db.Exec(deleteSql, id)

	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}
