package repository

import (
	"database/sql"
	"fmt"
	"transactions/models"

	"github.com/shopspring/decimal"
)

type TransactionRepositoryInterface interface {
	SubmitTransaction(sourceID, destID int64, amount models.Money) error
}

type TransactionRepository struct {
	DB *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{DB: db}
}

func (r *TransactionRepository) SubmitTransaction(sourceID, destID int64, amount models.Money) error {
	tx, err := r.DB.Begin()
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
