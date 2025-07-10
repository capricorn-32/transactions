package handler

import (
	"encoding/json"
	"net/http"
	"transactions/service"
)

type TransactionHandler struct {
	Service service.TransactionServiceInterface
}

func NewTransactionHandler(service service.TransactionServiceInterface) *TransactionHandler {
	return &TransactionHandler{Service: service}
}

func (h *TransactionHandler) SubmitTransaction(w http.ResponseWriter, r *http.Request) {
	var req struct {
		SourceAccountID      int64  `json:"source_account_id"`
		DestinationAccountID int64  `json:"destination_account_id"`
		Amount               string `json:"amount"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "invalid request",
		})
		return
	}
	if err := h.Service.SubmitTransaction(req.SourceAccountID, req.DestinationAccountID, req.Amount); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "transaction submitted successfully",
	})
}
