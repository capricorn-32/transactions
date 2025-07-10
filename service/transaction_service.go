package service

import (
	"database/sql"
	"fmt"
	"transactions/models"
	"transactions/repository"

	"github.com/shopspring/decimal"
)

// Define the interface here for use in both service and handler
// TransactionServiceInterface defines the contract for transaction service
// so it can be used for both real and mock implementations.
type TransactionServiceInterface interface {
	SubmitTransaction(sourceID, destID int64, amount models.Money) error
}

type TransactionService struct {
	DB   *sql.DB
	Repo *repository.TransactionRepository
}

func NewTransactionService(db *sql.DB, repo *repository.TransactionRepository) *TransactionService {
	return &TransactionService{DB: db, Repo: repo}
}

func (s *TransactionService) SubmitTransaction(sourceID, destID int64, amount models.Money) error {
	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	amt := amount.Decimal
	if amt.LessThanOrEqual(decimal.Zero) {
		return fmt.Errorf("amount must be positive")
	}

	// Check source balance
	var sourceBalanceStr string
	err = tx.QueryRow("SELECT balance FROM accounts WHERE account_id = $1 FOR UPDATE", sourceID).Scan(&sourceBalanceStr)
	if err != nil {
		return fmt.Errorf("source account not found or error: %w", err)
	}
	sourceBalance, err := decimal.NewFromString(sourceBalanceStr)
	if err != nil {
		return fmt.Errorf("invalid source account balance: %w", err)
	}
	if sourceBalance.LessThan(amt) {
		return fmt.Errorf("insufficient funds")
	}

	// Deduct from source
	newSourceBalance := sourceBalance.Sub(amt)
	_, err = tx.Exec("UPDATE accounts SET balance = $1 WHERE account_id = $2", newSourceBalance.String(), sourceID)
	if err != nil {
		return err
	}

	// Add to destination
	var destBalanceStr string
	err = tx.QueryRow("SELECT balance FROM accounts WHERE account_id = $1 FOR UPDATE", destID).Scan(&destBalanceStr)
	if err != nil {
		return fmt.Errorf("destination account not found or error: %w", err)
	}
	destBalance, err := decimal.NewFromString(destBalanceStr)
	if err != nil {
		return fmt.Errorf("invalid destination account balance: %w", err)
	}
	newDestBalance := destBalance.Add(amt)
	_, err = tx.Exec("UPDATE accounts SET balance = $1 WHERE account_id = $2", newDestBalance.String(), destID)
	if err != nil {
		return err
	}

	// Log transaction
	_, err = tx.Exec(`INSERT INTO transactions (source_account_id, destination_account_id, amount) VALUES ($1, $2, $3)`, sourceID, destID, amount.String())
	if err != nil {
		return err
	}

	// Commit transaction
	return tx.Commit()
}

var _ TransactionServiceInterface = (*TransactionService)(nil)
