package config

import (
	"log"

	"github.com/fahmiabd/go-order-api/internal/model"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	if db == nil {
		log.Println("skip migration: db not connected")
		return
	}

	err := db.AutoMigrate(
		&model.User{},
		&model.Order{},
	)

	if err != nil {
		log.Println("migration failed:", err)
	} else {
		log.Println("migration success")
	}
}
