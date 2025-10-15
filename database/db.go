package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "expense"
)

func Connect() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		return nil, fmt.Errorf("failed to create connection: %v", err)
	}

	query := `
	CREATE TABLE IF NOT EXISTS ledger (
		record_id SERIAL PRIMARY KEY,
		record_type TEXT,
		amount DECIMAL(10, 2),
		currency VARCHAR(3),
		description TEXT,
		date DATE,
		category TEXT
	);
	`

	_, err = db.Exec(query)

	if err != nil {
		return nil, fmt.Errorf("failed to create tables: %v", err)
	}

	return db, nil

}
