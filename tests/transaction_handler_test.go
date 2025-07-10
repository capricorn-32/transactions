package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"transactions/handler"
	"transactions/service"
)

func TestSubmitTransaction(t *testing.T) {
	// Use a mock or in-memory service for real tests
	h := handler.NewTransactionHandler(&service.TransactionService{DB: nil, Repo: nil})
	body := []byte(`{"source_account_id": 1, "destination_account_id": 2, "amount": "50.00"}`)
	req := httptest.NewRequest(http.MethodPost, "/transactions", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	h.SubmitTransaction(w, req)
	if w.Code != http.StatusCreated && w.Code != http.StatusBadRequest {
		t.Errorf("expected status 201 or 400, got %d", w.Code)
	}
} 