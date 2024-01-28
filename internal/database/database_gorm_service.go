package database

import (
	"gorm.io/gorm"
)

type DatabaseGormService interface {
	InitDB()
	CloseDB() error
	Migrate()

	GetDB() *gorm.DB
}
