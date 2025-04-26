// filepath: d:\study-project\dibimbing_golang\dibimbing_golang_capstone\main.go
package main

import (
	"log"
	"dibimbing_golang_capstone/config"
	"dibimbing_golang_capstone/middleware"
	"dibimbing_golang_capstone/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Inisialisasi database dan auto migration
	db := config.InitDB()
	
	// Optional: Tambahkan pengecekan koneksi database
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get database instance:", err)
	}
	defer sqlDB.Close()

	// Inisialisasi router
	r := gin.Default()

	// Middleware
	r.Use(middleware.CORSMiddleware())

	// Setup routes
	routes.SetupRoutes(r)

	log.Println("Server is running on port :8081")
	if err := r.Run(":8081"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}