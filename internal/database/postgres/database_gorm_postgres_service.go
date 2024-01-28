package postgres

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"task-tracker-backend/internal/model"
	"task-tracker-backend/internal/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseGormPostgresService struct {
	db *gorm.DB
}

var service *DatabaseGormPostgresService
var serviceOnce sync.Once

// Get singleton instance of Database service
func GetService() *DatabaseGormPostgresService {
	serviceOnce.Do(func() {
		service = &DatabaseGormPostgresService{}
		service.InitDB()
		service.Migrate()
	})

	return service
}

// Initialise connection to postgres database.
// Properties are taken from .env file
//
// Called by GetService() once, when creating Service
func (dbService *DatabaseGormPostgresService) InitDB() {
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	utils.HandleError(err)

	connectionProperties := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		port,
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))

	dbService.db = dbService.connectToPostgress(connectionProperties)
}

// A way to close database connection.
//
// It is optional to call this method, because connection
// will close automatically after program stops
func (dbService *DatabaseGormPostgresService) CloseDB() error {
	db, err := dbService.db.DB()
	utils.HandleError(err)

	return db.Close()
}

// Migrate database to current schema. Has list of all model structs
//
// Called by GetService() once, when creating Service
func (dbService *DatabaseGormPostgresService) Migrate() {
	dbService.db.AutoMigrate(&model.User{}, &model.Task{})
}

// Getter for DB field. DB is a database connection to perform all queries
func (dbService *DatabaseGormPostgresService) GetDB() *gorm.DB {
	return dbService.db
}

// Open session with postgres database. Call log.Fatal() on error
func (dbService *DatabaseGormPostgresService) connectToPostgress(connectionProperties string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(connectionProperties), &gorm.Config{})
	utils.HandleError(err)

	return db
}
