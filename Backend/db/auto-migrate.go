package db

import (
	"github.com/riz254/Go-Ticketing-System.git/models"
	"gorm.io/gorm"
)

func DBMigrator(db *gorm.DB) error {
	return db.AutoMigrate(&models.Event{})
}
