package handler

import (
	"encoding/json"
	"net/http"
	"transactions/models"
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
		SourceAccountID      int64        `json:"source_account_id"`
		DestinationAccountID int64        `json:"destination_account_id"`
		Amount               models.Money `json:"amount"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "invalid request: could not decode JSON")
		return
	}

	// Input validation
	if req.SourceAccountID <= 0 || req.DestinationAccountID <= 0 {
		WriteErrorResponse(w, http.StatusBadRequest, "source_account_id and destination_account_id must be positive integers")
		return
	}

	// Check if source and destination accounts are different
	if req.SourceAccountID == req.DestinationAccountID {
		WriteErrorResponse(w, http.StatusBadRequest, "source_account_id and destination_account_id must not be the same")
		return
	}

	// Check if amount is a valid positive number
	if req.Amount.Decimal.LessThanOrEqual(models.Money{}.Decimal) {
		WriteErrorResponse(w, http.StatusBadRequest, "amount must be a valid positive number")
		return
	}

	// Log the error for debugging purposes
	if err := h.Service.SubmitTransaction(req.SourceAccountID, req.DestinationAccountID, req.Amount); err != nil {
		WriteErrorResponse(w, http.StatusInternalServerError, "failed to submit transaction: "+err.Error())
		return
	}

	// If everything is successful, return a success response
	WriteCreatedResponse(w, "Transaction submitted successfully")
}
