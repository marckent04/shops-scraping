package database

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect(dsn string) {
	conn, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	log.Println("Database connected successfully")

	db = conn

	runMigrations()
}

func runMigrations() {
	db.AutoMigrate(&CartLine{})
}
