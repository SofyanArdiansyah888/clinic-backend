package config

import (
	"backend/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	// OPEN CONNECTION
	//@TODO: Uncomment the following line to use environment variable for database connection
	dsn := os.Getenv("DATABASE_URL")
	// dsn := "host=localhost user=root password=#Bismillah_01 dbname=clinic port=5433 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// AUTOMIGRATE
	err = db.AutoMigrate(models.GetModels()...)
	// err = db.AutoMigrate(&models.Barang{}, &models.StokMovement{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	DB = db
	log.Println("Database connected")
	return DB
}
