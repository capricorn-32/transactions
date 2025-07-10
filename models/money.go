package models

import (
	"encoding/json"
	"fmt"

	"github.com/shopspring/decimal"
)

type Money struct {
	decimal.Decimal
}

func (m *Money) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	d, err := decimal.NewFromString(s)
	if err != nil {
		return fmt.Errorf("invalid money value: %w", err)
	}
	m.Decimal = d
	return nil
}

func (m Money) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.String())
}

func NewMoneyFromString(s string) (Money, error) {
	d, err := decimal.NewFromString(s)
	if err != nil {
		return Money{}, err
	}
	return Money{d}, nil
}
