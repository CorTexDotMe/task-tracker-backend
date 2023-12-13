package postgres

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func ConnectToPostgress(connectionProperties string) *sql.DB {
	db, err := sql.Open("postgres", connectionProperties)
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}

	return db
}
