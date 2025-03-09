package internalsql

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Connect(datasSourceName string) (*sql.DB, error) {
	db, err := sql.Open("mysql", datasSourceName)
	if err != nil {
		log.Fatalf("Error on initializing database connection: %v", err)
		return nil, err
	}

	return db, nil
}
