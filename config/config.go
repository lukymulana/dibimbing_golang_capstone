package config

import (
	"log"
	"dibimbing_golang_capstone/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/travel_booking?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto migrate entities
	err = db.AutoMigrate(
		&entity.User{},
		&entity.Trip{},
		&entity.Booking{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Database migrated successfully")
	DB = db
	return db
}