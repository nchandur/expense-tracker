package models

import (
	"fmt"
	"strings"
	"time"
)

type RecordType string

const (
	Expense RecordType = "expense"
	Income  RecordType = "income"
)

type Record struct {
	ID          int `json:"record_id"`
	RecordType  `json:"record_type"`
	Amount      float64   `json:"amount"`
	Currency    string    `json:"currency"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	Category    string    `json:"category"`
}

func NewRecord(recordType RecordType, amount float64, currency string, description string, date time.Time, category string) (Record, error) {
	if (recordType != Expense) && (recordType != Income) {
		return Record{}, fmt.Errorf("invalid record type: expense/income only")
	}

	if amount < 0 {
		return Record{}, fmt.Errorf("invalid value for amount. must be non-negative")
	}

	if len(currency) > 3 {
		currency = strings.ToUpper(currency[:3])
	}

	if category == "" {
		category = "miscellaneous"
	}

	return Record{
		RecordType:  recordType,
		Amount:      amount,
		Currency:    currency,
		Description: description,
		Date:        date,
		Category:    strings.TrimSpace(strings.ToLower(category)),
	}, nil

}
