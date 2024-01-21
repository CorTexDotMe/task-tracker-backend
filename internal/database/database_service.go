package database

import (
	"fmt"
	"os"
	"strconv"
	"task-tracker-backend/internal/database/postgres"
	"task-tracker-backend/internal/model"
	"task-tracker-backend/internal/utils"

	"gorm.io/gorm"
)

// Database session to run all queries
var DB *gorm.DB

// Initialise connection to postgres database.
// Properties are taken from .env file
func InitDB() {
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	utils.HandleError(err)

	connectionProperties := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		port,
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))

	DB = postgres.ConnectToPostgress(connectionProperties)
}

// A way to close database connection.
//
// It is optional to call this method, because connection
// will close automatically after program stops
func CloseDB() error {
	db, err := DB.DB()
	utils.HandleError(err)

	return db.Close()
}

// Migrate database to current schema. Has list of all model structs
func Migrate() {
	DB.AutoMigrate(&model.User{}, &model.Task{})
}
