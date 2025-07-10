package handler

import (
	"encoding/json"
	"fmt"
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
			"error":   "invalid request: could not decode JSON",
		})
		return
	}
	// Input validation
	if req.SourceAccountID <= 0 || req.DestinationAccountID <= 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "source_account_id and destination_account_id must be positive integers",
		})
		return
	}
	if req.SourceAccountID == req.DestinationAccountID {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "source_account_id and destination_account_id must not be the same",
		})
		return
	}
	if req.Amount == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "amount is required",
		})
		return
	}
	var amt float64
	if _, err := fmt.Sscanf(req.Amount, "%f", &amt); err != nil || amt <= 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "amount must be a valid positive number",
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
