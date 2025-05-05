package config

import (
	"backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	// OPEN CONNECTION
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// AUTOMIGRATE
	err = db.AutoMigrate(models.GetModels()...)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	DB = db
	log.Println("Database connected")
	return DB
}
