package repository

import (
	"database/sql"
	"transactions/models"
)

type TransactionRepository struct {
	DB *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{DB: db}
}

func (r *TransactionRepository) SubmitTransaction(sourceID, destID int64, amount models.Money) error {
	_, err := r.DB.Exec(`INSERT INTO transactions (source_account_id, destination_account_id, amount) VALUES ($1, $2, $3)`, sourceID, destID, amount.String())
	return err
}
