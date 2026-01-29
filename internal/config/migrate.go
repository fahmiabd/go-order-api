package config

import (
	"log"

	"github.com/fahmiabd/go-order-api/internal/models"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	if db == nil {
		log.Println("skip migration: db not connected")
		return
	}

	err := db.AutoMigrate(
		&models.User{},
		&models.Order{},
		&models.Product{},
	)

	if err != nil {
		log.Println("migration failed:", err)
	} else {
		log.Println("migration success")
	}
}
