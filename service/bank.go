package service

type BankResponse struct {
	BankID int    `json:"bank_id"`
	Name   string `json:"name"`
	Status int    `json:"status"`
}

type BankRequest struct {
	BankID int    `json:"bank_id"`
	Name   string `json:"name"`
	Status int    `json:"status"`
}

type BankService interface {
	GetBanksAll() ([]BankResponse, error)
	GetBank(int) (*BankResponse, error)
	PostBank(BankRequest) (*BankResponse, error)
	PutBank(BankRequest) (*BankResponse, error)
	DeleteBank(int) (*BankResponse, error)
}
