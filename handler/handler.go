package handler

import (
	"transactions/service"
)

type Handler struct {
	Account     *AccountHandler
	Transaction *TransactionHandler
}

func NewHandler(accountService *service.AccountService, transactionService *service.TransactionService) *Handler {
	return &Handler{
		Account:     NewAccountHandler(accountService),
		Transaction: NewTransactionHandler(transactionService),
	}
}
