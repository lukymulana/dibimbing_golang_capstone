// filepath: d:\study-project\dibimbing_golang\dibimbing_golang_capstone\main.go
package main

import (
    "log"

    "github.com/gin-gonic/gin"
    "dibimbing_golang_capstone/config"
    "dibimbing_golang_capstone/middleware"
)

func main() {
    // Inisialisasi database
    config.InitDB()

    // Inisialisasi router
    r := gin.Default()

    // Middleware
    r.Use(middleware.CORSMiddleware())

    // Setup routes
    setupRoutes(r)

    log.Println("Server is running on port :8080")
    if err := r.Run(":8080"); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}