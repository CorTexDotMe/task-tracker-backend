package database

import (
	"database/sql"
	"fmt"
	"log"
	"task-tracker-backend/internal/database/postgres"
	"task-tracker-backend/internal/utils"

	"github.com/golang-migrate/migrate/v4"
	migratePostgres "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// TODO move vars into .env file
const (
	HOST                 = "host.docker.internal"
	PORT                 = 5432
	USER                 = "postgres"
	PASSWORD             = "password"
	DATABASE_NAME        = "task-tracker"
	MIGRATION_FILES_PATH = "file://../../internal/database/migrations"
	DATABASE_DRIVER_NAME = "postgres"
)

var DATABASE_CONNECTION *sql.DB

func InitDB() {
	connectionProperties := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		HOST, PORT, USER, PASSWORD, DATABASE_NAME)

	DATABASE_CONNECTION = postgres.ConnectToPostgress(connectionProperties)
}

func CloseDB() error {
	return DATABASE_CONNECTION.Close()
}

func Migrate() {
	utils.HandleError(DATABASE_CONNECTION.Ping())

	driver, _ := migratePostgres.WithInstance(DATABASE_CONNECTION, &migratePostgres.Config{})
	migratoin, migrateErr := migrate.NewWithDatabaseInstance(
		MIGRATION_FILES_PATH, DATABASE_DRIVER_NAME, driver)
	utils.HandleError(migrateErr)

	if err := migratoin.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
}
