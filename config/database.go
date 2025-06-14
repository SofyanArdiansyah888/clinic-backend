package config

import (
	"backend/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	// OPEN CONNECTION
	//@TODO: Uncomment the following line to use environment variable for database connection
	//dsn := os.Getenv("DATABASE_URL")
	dsn := "u274107390_schaleup:#Bismillah_01@tcp(153.92.11.7:3306)/u274107390_schaleup?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info),
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
