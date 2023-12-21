package database

import (
	"fmt"
	"task-tracker-backend/internal/database/postgres"
	model "task-tracker-backend/internal/models"
	"task-tracker-backend/internal/utils"

	"gorm.io/gorm"
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

var DB *gorm.DB

func InitDB() {
	connectionProperties := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		HOST, PORT, USER, PASSWORD, DATABASE_NAME)

	DB = postgres.ConnectToPostgress(connectionProperties)
}

func CloseDB() error {
	db, err := DB.DB()
	utils.HandleError(err)

	return db.Close()
}

func Migrate() {
	DB.AutoMigrate(&model.User{})
}
