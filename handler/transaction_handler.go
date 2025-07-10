package handler

import (
	"encoding/json"
	"net/http"
	"transactions/service"
)

type TransactionHandler struct {
	Service *service.TransactionService
}

func NewTransactionHandler(service *service.TransactionService) *TransactionHandler {
	return &TransactionHandler{Service: service}
}

func (h *TransactionHandler) SubmitTransaction(w http.ResponseWriter, r *http.Request) {
	var req struct {
		SourceAccountID      int64  `json:"source_account_id"`
		DestinationAccountID int64  `json:"destination_account_id"`
		Amount               string `json:"amount"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	if err := h.Service.SubmitTransaction(req.SourceAccountID, req.DestinationAccountID, req.Amount); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
} 