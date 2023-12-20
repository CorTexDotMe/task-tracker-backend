package postgres

import (
	"task-tracker-backend/internal/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToPostgress(connectionProperties string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(connectionProperties), &gorm.Config{})
	utils.HandleError(err)

	return db
}
