package tests

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"transactions/handler"
)

type fakeTransactionService struct{}

func (f *fakeTransactionService) SubmitTransaction(sourceID, destID int64, amount string) error {
	if sourceID == 0 || destID == 0 {
		return errors.New("invalid account id")
	}
	if amount == "0" {
		return errors.New("amount must be positive")
	}
	if amount == "9999" {
		return errors.New("insufficient funds")
	}
	return nil
}

func newTestTransactionHandler() *handler.TransactionHandler {
	return handler.NewTransactionHandler(&fakeTransactionService{})
}

func TestSubmitTransaction_Success(t *testing.T) {
	h := newTestTransactionHandler()
	body := []byte(`{"source_account_id": 1, "destination_account_id": 2, "amount": "100.00"}`)
	req := httptest.NewRequest(http.MethodPost, "/transactions", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	h.SubmitTransaction(w, req)
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
	if resp["message"] != "transaction submitted successfully" {
		t.Errorf("unexpected message: %v", resp["message"])
	}
}

func TestSubmitTransaction_Error(t *testing.T) {
	h := newTestTransactionHandler()
	body := []byte(`{"source_account_id": 0, "destination_account_id": 2, "amount": "50.00"}`)
	req := httptest.NewRequest(http.MethodPost, "/transactions", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	h.SubmitTransaction(w, req)
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
