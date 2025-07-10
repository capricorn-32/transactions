package service

import (
	"database/sql"
	"fmt"
	"transactions/repository"
)

// Define the interface here for use in both service and handler
// TransactionServiceInterface defines the contract for transaction service
// so it can be used for both real and mock implementations.
type TransactionServiceInterface interface {
	SubmitTransaction(sourceID, destID int64, amount string) error
}

type TransactionService struct {
	DB   *sql.DB
	Repo *repository.TransactionRepository
}

func NewTransactionService(db *sql.DB, repo *repository.TransactionRepository) *TransactionService {
	return &TransactionService{DB: db, Repo: repo}
}

func (s *TransactionService) SubmitTransaction(sourceID, destID int64, amount string) error {
	// Start transaction
	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	// Check source balance
	var sourceBalance string
	err = tx.QueryRow("SELECT balance FROM accounts WHERE account_id = $1 FOR UPDATE", sourceID).Scan(&sourceBalance)
	if err != nil {
		return fmt.Errorf("source account not found or error: %w", err)
	}

	// Check sufficient funds
	var src, amt float64
	fmt.Sscanf(sourceBalance, "%f", &src)
	fmt.Sscanf(amount, "%f", &amt)
	if src < amt {
		return fmt.Errorf("insufficient funds")
	}

	// Deduct from source
	_, err = tx.Exec("UPDATE accounts SET balance = balance - $1 WHERE account_id = $2", amount, sourceID)
	if err != nil {
		return err
	}
	// Add to destination
	_, err = tx.Exec("UPDATE accounts SET balance = balance + $1 WHERE account_id = $2", amount, destID)
	if err != nil {
		return err
	}
	// Log transaction
	_, err = tx.Exec(`INSERT INTO transactions (source_account_id, destination_account_id, amount) VALUES ($1, $2, $3)`, sourceID, destID, amount)
	if err != nil {
		return err
	}
	return nil
}

var _ TransactionServiceInterface = (*TransactionService)(nil)
