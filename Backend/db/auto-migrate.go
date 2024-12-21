package db

import (
	"github.com/riz254/Go-Ticketing-System.git/models"
	"gorm.io/gorm"
)

func DBMigrator(db *gorm.DB) error {
	// Drop tables if they exist
	err := db.Migrator().DropTable(&models.Event{}, &models.Ticket{}, &models.User{})
	if err != nil {
		return err
	}

	// Recreate tables
	err = db.AutoMigrate(&models.Event{}, &models.Ticket{}, &models.User{})
	if err != nil {
		return err
	}

	return nil
}
