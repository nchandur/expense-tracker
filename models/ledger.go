package models

import "time"

type Ledger []Record

func NewLedger() Ledger {
	return make(Ledger, 0)
}

type LedgerService interface {
	CreateRecord(record Record) error
	RetrieveRecords(start, end time.Time) (Ledger, error)
	UpdateRecord(id int, record Record) error
	DeleteRecord(id int) error
}