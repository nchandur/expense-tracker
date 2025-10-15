package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/nchandur/expense-tracker/models"
)

type LedgerService struct {
	db *sql.DB
}

func NewLedgerService(db *sql.DB) (LedgerService, error) {
	err := db.Ping()

	if err != nil {
		return LedgerService{nil}, err
	}

	return LedgerService{db}, nil
}

func (l LedgerService) CreateRecord(record models.Record) error {
	query := `
	INSERT INTO ledger (record_type, amount, currency, description, date, category) VALUES ($1, $2, $3, $4, $5, $6);
	`

	_, err := l.db.Exec(query, record.RecordType, record.Amount, record.Currency, record.Description, record.Date, record.Category)

	if err != nil {
		return fmt.Errorf("failed to create record: %v", err)
	}

	return nil
}

func (l LedgerService) RetrieveRecords(start, end time.Time) (models.Ledger, error) {

	if start.After(end) {
		return nil, fmt.Errorf("start date must precede end date")
	}

	ledger := models.NewLedger()

	query := `
	SELECT * FROM ledger WhERE date BETWEEN $1 AND $2;
	`

	rows, err := l.db.Query(query, start, end)

	if err != nil {
		return nil, fmt.Errorf("failed to retrieve records: %v", err)
	}

	for rows.Next() {
		var record models.Record

		rows.Scan(&record.ID, &record.RecordType, &record.Amount, &record.Currency, &record.Description, &record.Date, &record.Category)

		ledger = append(ledger, record)

	}

	return ledger, nil
}

func (l LedgerService) DeleteRecord(id int) error {

	query := `
	DELETE FROM ledger WHERE record_id = $1;
	`

	res, err := l.db.Exec(query, id)

	if err != nil {
		return fmt.Errorf("failed to delete record with ID: %d: %v", id, err)
	}

	affected, _ := res.RowsAffected()

	if affected == 0 {
		return fmt.Errorf("no record with ID: %d found", id)
	}

	return nil

}

func (l LedgerService) UpdateRecord(id int, record models.Record) error {
	err := l.DeleteRecord(id)

	if err != nil {
		return fmt.Errorf("failed to update record with ID: %d: %v", id, err)
	}

	err = l.CreateRecord(record)

	if err != nil {
		return fmt.Errorf("failed to update record: %v", err)
	}

	return err

}
