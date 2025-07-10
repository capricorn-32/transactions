package tests

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"transactions/handler"
	"transactions/models"
	"transactions/service"

	"github.com/gorilla/mux"
)

type mockAccountRepo struct{}

func (m *mockAccountRepo) CreateAccount(accountID int64, initialBalance string) error {
	if accountID == 999 {
		return errors.New("duplicate account")
	}
	return nil
}
func (m *mockAccountRepo) GetAccount(accountID int64) (*models.Account, error) {
	return &models.Account{AccountID: accountID, Balance: "100.00"}, nil
}

func TestCreateAccount_Success(t *testing.T) {
	h := handler.NewAccountHandler(&service.AccountService{Repo: &mockAccountRepo{}})
	body := []byte(`{"account_id": 1, "initial_balance": "100.00"}`)
	req := httptest.NewRequest(http.MethodPost, "/accounts", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	h.CreateAccount(w, req)
	if w.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", w.Code)
	}
	var resp map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("invalid JSON response: %v", err)
	}
	if !resp["success"].(bool) {
		t.Errorf("expected success true, got false")
	}
	if resp["message"] != "account created successfully" {
		t.Errorf("unexpected message: %v", resp["message"])
	}
}

func TestCreateAccount_Error(t *testing.T) {
	h := handler.NewAccountHandler(&service.AccountService{Repo: &mockAccountRepo{}})
	body := []byte(`{"account_id": 999, "initial_balance": "100.00"}`)
	req := httptest.NewRequest(http.MethodPost, "/accounts", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	h.CreateAccount(w, req)
	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", w.Code)
	}
	var resp map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("invalid JSON response: %v", err)
	}
	if resp["success"].(bool) {
		t.Errorf("expected success false, got true")
	}
	if resp["error"] == nil || resp["error"] == "" {
		t.Errorf("expected error message, got: %v", resp["error"])
	}
}

func TestGetAccount_Success(t *testing.T) {
	h := handler.NewAccountHandler(&service.AccountService{Repo: &mockAccountRepo{}})
	req := httptest.NewRequest(http.MethodGet, "/accounts/1", nil)
	req = mux.SetURLVars(req, map[string]string{"account_id": "1"})
	w := httptest.NewRecorder()

	h.GetAccount(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}
	var resp struct {
		Success bool            `json:"success"`
		Data    *models.Account `json:"data"`
	}
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("invalid JSON response: %v", err)
	}
	if !resp.Success {
		t.Errorf("expected success true, got false")
	}
	if resp.Data == nil || resp.Data.AccountID != 1 {
		t.Errorf("unexpected data: %+v", resp.Data)
	}
}

func TestGetAccount_BadRequest(t *testing.T) {
	h := handler.NewAccountHandler(&service.AccountService{Repo: &mockAccountRepo{}})
	req := httptest.NewRequest(http.MethodGet, "/accounts/abc", nil)
	req = mux.SetURLVars(req, map[string]string{"account_id": "abc"})
	w := httptest.NewRecorder()

	h.GetAccount(w, req)
	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", w.Code)
	}
	var resp map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("invalid JSON response: %v", err)
	}
	if resp["success"].(bool) {
		t.Errorf("expected success false, got true")
	}
	if resp["error"] == nil || resp["error"] == "" {
		t.Errorf("expected error message, got: %v", resp["error"])
	}
}
