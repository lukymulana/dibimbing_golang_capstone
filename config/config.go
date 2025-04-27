package config

import (
	"log"
	"dibimbing_golang_capstone/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func buildDSN() string {
	dsn := GetEnv("DB_DSN", "")
	if dsn != "" {
		return dsn
	}
	host := GetEnv("DB_HOST", "localhost")
	port := GetEnv("DB_PORT", "3306")
	user := GetEnv("DB_USER", "root")
	password := GetEnv("DB_PASSWORD", "")
	dbname := GetEnv("DB_NAME", "travel_booking")
	return user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
}

func InitDB() *gorm.DB {
	LoadEnv()
	dsn := buildDSN()
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