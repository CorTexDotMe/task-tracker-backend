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

var DB *gorm.DB

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

func CloseDB() error {
	db, err := DB.DB()
	utils.HandleError(err)

	return db.Close()
}

func Migrate() {
	DB.AutoMigrate(&model.User{}, &model.Task{})
}
