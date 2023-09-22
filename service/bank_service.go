package service

import (
	"database/sql"
	"log"
	"transaction/repository"
)

type bankService struct {
	bankRepo repository.BankRepository
}

func NewBankService(bankRepo repository.BankRepository) BankService {
	return bankService{bankRepo: bankRepo}
}

func (s bankService) GetBanksAll() ([]BankResponse, error) {

	banks, err := s.bankRepo.GetBanksAll()
	if err != nil {
		log.Println("Error while getting banks:", err)
		return nil, err
	}

	bankResponses := []BankResponse{}
	for _, bank := range banks {
		bankResponse := BankResponse{
			BankID: bank.BankID,
			Name:   bank.Name,
			Status: bank.Status,
		}
		bankResponses = append(bankResponses, bankResponse)
	}

	return bankResponses, nil
}

func (s bankService) GetBank(id int) (*BankResponse, error) {
	bank, err := s.bankRepo.GetBankById(id)

	if err != nil {

		if err == sql.ErrNoRows {
			return nil, err
		}

		log.Println("Error while getting bank:", err)
		return nil, err
	}

	bankResponse := BankResponse{
		BankID: bank.BankID,
		Name:   bank.Name,
		Status: bank.Status,
	}

	return &bankResponse, nil
}

func (s bankService) PostBank(bankReq BankRequest) (*BankResponse, error) {
	bank := repository.BankModel{
		Name:   bankReq.Name,
		Status: bankReq.Status,
	}

	newBank, err := s.bankRepo.PostBank(bank)

	if err != nil {
		log.Println("Error while posting bank:", err)
		return nil, err
	}
	bankResponse := BankResponse{
		BankID: newBank.BankID,
		Name:   newBank.Name,
		Status: newBank.Status,
	}

	return &bankResponse, nil
}

func (s bankService) PutBank(bankReq BankRequest) (*BankResponse, error) {
	bank := repository.BankModel{
		BankID: bankReq.BankID,
		Name:   bankReq.Name,
		Status: bankReq.Status,
	}

	_, err := s.bankRepo.PutBank(bank)
	if err != nil {
		log.Println("Error while putting bank:", err)
		return nil, err
	}

	bankResponse := BankResponse{
		BankID: bank.BankID,
		Name:   bank.Name,
		Status: bank.Status,
	}

	return &bankResponse, nil
}

func (s bankService) DeleteBank(id int) (*BankResponse, error) {
	_, err := s.bankRepo.DeleteBank(id)
	if err != nil {
		log.Println("Error while deleting bank:", err)
		return nil, err
	}

	bankResponse := BankResponse{
		BankID: id,
	}

	return &bankResponse, nil
}
