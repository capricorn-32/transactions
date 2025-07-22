package service

import (
	"transactions/models"
	"transactions/repository"
)

type TransactionServiceInterface interface {
	SubmitTransaction(sourceID, destID int64, amount models.Money) error
}

type TransactionService struct {
	Repo repository.TransactionRepositoryInterface
}

func NewTransactionService(repo repository.TransactionRepositoryInterface) *TransactionService {
	return &TransactionService{Repo: repo}
}

func (s *TransactionService) SubmitTransaction(sourceID, destID int64, amount models.Money) error {
	return s.Repo.SubmitTransaction(sourceID, destID, amount)
}
