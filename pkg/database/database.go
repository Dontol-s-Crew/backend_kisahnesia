package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func GetDatabase(dbAddress string, dbUsername string, dbPassword string, dbName string) *sql.DB {
	log.Printf("INFO GetDatabase database connection: starting database connection process")

	dataSourceName := fmt.Sprintf("host=%s port=5432 user=%s password=%s dbname=%s sslmode=disable",
		dbAddress, dbUsername, dbPassword, dbName)
	fmt.Print(dataSourceName)
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatalf("Error GetDatabase sql open connection fatal error: %v", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("ERROR GetDatabase db ping fatal error: %v", err)
	}
	log.Printf("INFO GetDatabase database connectionn: established successfully with %s\n", dataSourceName)
	return db
}
