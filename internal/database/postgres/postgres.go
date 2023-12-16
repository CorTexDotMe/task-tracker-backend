package postgres

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// TODO take from .env
const DATABASE_DRIVER_NAME = "postgres"

func ConnectToPostgress(connectionProperties string) *sql.DB {
	db, err := sql.Open(DATABASE_DRIVER_NAME, connectionProperties)
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}

	return db
}
