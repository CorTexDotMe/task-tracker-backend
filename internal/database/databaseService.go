package database

import (
	"database/sql"
	"fmt"
	"task-tracker-backend/internal/database/postgres"
)

// TODO: move vars into .env file
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "task-tracker"
)

var DB *sql.DB

func InitDB() {
	connectionProperties := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	DB = postgres.ConnectToPostgress(connectionProperties)
}

func CloseDB() error {
	return DB.Close()
}
