package main

import (
	"log"
	"time"

	"github.com/nchandur/expense-tracker/database"
	"github.com/nchandur/expense-tracker/models"
)

func main() {
	db, err := database.Connect()

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	ledgerService, err := database.NewLedgerService(db)

	if err != nil {
		log.Fatal(err)
	}

	record, err := models.NewRecord("expense", 10.00, "USD", "Charger", time.Now(), "Electronics But Updated")

	if err != nil {
		log.Fatal(err)
	}

	err = ledgerService.UpdateRecord(2, record)

	if err != nil {
		log.Fatal(err)
	}

}
