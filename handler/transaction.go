package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"transaction/service"

	"github.com/gorilla/mux"
)

type transactionHandler struct {
	transactionSrv service.TransactionService
}

func NewTransactionHandler(transactionSrv service.TransactionService) transactionHandler {
	return transactionHandler{transactionSrv: transactionSrv}
}

func (h transactionHandler) GetTransactions(w http.ResponseWriter, r *http.Request) {
	transactions, err := h.transactionSrv.GetTransactionsAll()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transactions)
}

func (h transactionHandler) GetTransactionByCustomerId(w http.ResponseWriter, r *http.Request) {
	customerId, _ := strconv.Atoi(mux.Vars(r)["customer_id"])

	transactions, err := h.transactionSrv.GetTransactionByCustomerId(customerId)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transactions)
}

func (h transactionHandler) GetTransactionByBankId(w http.ResponseWriter, r *http.Request) {
	bankId, _ := strconv.Atoi(mux.Vars(r)["bank_id"])

	transactions, err := h.transactionSrv.GetTransactionByBankId(bankId)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transactions)
}

func (h transactionHandler) PostTransaction(w http.ResponseWriter, r *http.Request) {

	transactionReq := service.TransactionRequest{}
	json.NewDecoder(r.Body).Decode(&transactionReq)

	transaction, err := h.transactionSrv.PostTransaction(transactionReq)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transaction)
}

func (h transactionHandler) PutTransaction(w http.ResponseWriter, r *http.Request) {

	transactionReq := service.TransactionRequest{}
	json.NewDecoder(r.Body).Decode(&transactionReq)

	transaction, err := h.transactionSrv.PutTransaction(transactionReq)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transaction)
}

func (h transactionHandler) DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	transactionId, _ := strconv.Atoi(mux.Vars(r)["transaction_id"])

	transaction, err := h.transactionSrv.DeleteTransaction(transactionId)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transaction)
}
