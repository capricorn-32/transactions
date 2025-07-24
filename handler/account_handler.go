package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"transactions/service"

	"github.com/gorilla/mux"
)

type AccountHandler struct {
	Service *service.AccountService
}

func NewAccountHandler(service *service.AccountService) *AccountHandler {
	return &AccountHandler{Service: service}
}

func (h *AccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var req struct {
		AccountID      int64  `json:"account_id"`
		InitialBalance string `json:"initial_balance"`
	}

	// WriteErrorResponse is a convenience function for 400 Bad Request errors
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "invalid request: could not decode JSON")
		return
	}

	// Input validation
	if req.AccountID <= 0 {
		WriteErrorResponse(w, http.StatusBadRequest, "account_id must be a positive integer")
	}

	if req.InitialBalance == "" {
		WriteErrorResponse(w, http.StatusBadRequest, "initial_balance is required")
		return
	}

	var bal float64
	if _, err := fmt.Sscanf(req.InitialBalance, "%f", &bal); err != nil || bal < 0 {
		WriteErrorResponse(w, http.StatusBadRequest, "initial_balance must be a valid non-negative number")
		return
	}

	if err := h.Service.CreateAccount(req.AccountID, req.InitialBalance); err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "failed to create account: "+err.Error())
		return
	}
	// If everything is successful, return a success response
	WriteCreatedResponse(w, "Account created successfully")
}

func (h *AccountHandler) GetAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["account_id"]
	accountID, err := strconv.ParseInt(idStr, 10, 64)
	// WriteErrorResponse is a convenience function for 400 Bad Request errors
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "invalid account id: "+err.Error())
		return
	}

	acc, err := h.Service.GetAccount(accountID)
	if err != nil {
		WriteErrorResponse(w, http.StatusNotFound, "account not found: "+err.Error())
		return
	}

	WriteSuccessResponse(w, http.StatusOK, "Account retrieved successfully", acc)
}
