package postgres

import (
	"task-tracker-backend/internal/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Open session with postgres database. Call log.Fatal() on error
//
// Uses gorm
func ConnectToPostgress(connectionProperties string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(connectionProperties), &gorm.Config{})
	utils.HandleError(err)

	return db
}
