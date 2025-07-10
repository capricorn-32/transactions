package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"transactions/handler"
	"transactions/models"
	"transactions/service"
)

type mockAccountRepo struct{}

func (m *mockAccountRepo) CreateAccount(accountID int64, initialBalance string) error {
	return nil // or return an error to test error handling
}
func (m *mockAccountRepo) GetAccount(accountID int64) (*models.Account, error) {
	return &models.Account{AccountID: accountID, Balance: "100.00"}, nil
}

func TestCreateAccount(t *testing.T) {
	// Use a mock or in-memory service for real tests
	h := handler.NewAccountHandler(&service.AccountService{Repo: &mockAccountRepo{}})
	body := []byte(`{"account_id": 1, "initial_balance": "100.00"}`)
	req := httptest.NewRequest(http.MethodPost, "/accounts", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	h.CreateAccount(w, req)
	if w.Code != http.StatusCreated && w.Code != http.StatusBadRequest {
		t.Errorf("expected status 201 or 400, got %d", w.Code)
	}
}
