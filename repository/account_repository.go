package repository

import (
	"database/sql"
	"transactions/models"
)

type AccountRepositoryInterface interface {
	CreateAccount(accountID int64, initialBalance string) error
	GetAccount(accountID int64) (*models.Account, error)
}

type AccountRepository struct {
	DB *sql.DB
}

func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{DB: db}
}

func (r *AccountRepository) CreateAccount(accountID int64, initialBalance string) error {
	_, err := r.DB.Exec("INSERT INTO accounts (account_id, balance) VALUES ($1, $2)", accountID, initialBalance)
	return err
}

func (r *AccountRepository) GetAccount(accountID int64) (*models.Account, error) {
	row := r.DB.QueryRow("SELECT account_id, balance FROM accounts WHERE account_id = $1", accountID)
	var acc models.Account
	if err := row.Scan(&acc.AccountID, &acc.Balance); err != nil {
		return nil, err
	}
	return &acc, nil
}
