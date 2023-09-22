package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"transaction/service"

	"github.com/gorilla/mux"
)

type bankHandler struct {
	bankSrv service.BankService
}

func NewBankHandler(bankSrv service.BankService) bankHandler {
	return bankHandler{bankSrv: bankSrv}
}

func (h bankHandler) GetBanks(w http.ResponseWriter, r *http.Request) {
	banks, err := h.bankSrv.GetBanksAll()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(banks)
}

func (h bankHandler) GetBank(w http.ResponseWriter, r *http.Request) {
	bankID, _ := strconv.Atoi(mux.Vars(r)["bank_id"])

	bank, err := h.bankSrv.GetBank(bankID)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bank)
}

func (h bankHandler) PostBank(w http.ResponseWriter, r *http.Request) {

	bankReq := service.BankRequest{}
	json.NewDecoder(r.Body).Decode(&bankReq)

	bank, err := h.bankSrv.PostBank(bankReq)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bank)
}

func (h bankHandler) PutBank(w http.ResponseWriter, r *http.Request) {

	bankReq := service.BankRequest{}
	json.NewDecoder(r.Body).Decode(&bankReq)

	bank, err := h.bankSrv.PutBank(bankReq)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bank)
}

func (h bankHandler) DeleteBank(w http.ResponseWriter, r *http.Request) {
	bankID, _ := strconv.Atoi(mux.Vars(r)["bank_id"])

	bank, err := h.bankSrv.DeleteBank(bankID)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bank)
}
