package database

import (
	"database/sql"
	"fmt"
	"log"
	"task-tracker-backend/internal/database/postgres"

	"github.com/golang-migrate/migrate/v4"
	migratePostgres "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// TODO: move vars into .env file
const (
	host     = "host.docker.internal"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "task-tracker"
)

var DATABASE_CONNECTION *sql.DB

func InitDB() {
	connectionProperties := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	DATABASE_CONNECTION = postgres.ConnectToPostgress(connectionProperties)
}

func CloseDB() error {
	return DATABASE_CONNECTION.Close()
}

func Migrate() {
	if err := DATABASE_CONNECTION.Ping(); err != nil {
		log.Fatal(err)
	}
	driver, _ := migratePostgres.WithInstance(DATABASE_CONNECTION, &migratePostgres.Config{})
	migratoin, migrateErr := migrate.NewWithDatabaseInstance(
		"file://../../internal/database/migrations", "postgres", driver)
	if migrateErr != nil {
		log.Fatal(migrateErr)
	}

	if err := migratoin.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
}
