package service

import (
	"transactions/repository"
)

type AccountService struct {
	Repo repository.AccountRepositoryInterface
}

func NewAccountService(repo repository.AccountRepositoryInterface) *AccountService {
	return &AccountService{Repo: repo}
}

func (s *AccountService) CreateAccount(accountID int64, initialBalance string) error {
	return s.Repo.CreateAccount(accountID, initialBalance)
}

func (s *AccountService) GetAccount(accountID int64) (interface{}, error) {
	return s.Repo.GetAccount(accountID)
}
