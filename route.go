// filepath: d:\study-project\dibimbing_golang\dibimbing_golang_capstone\route.go
package main

import (
    "dibimbing_golang_capstone/controller"
    "dibimbing_golang_capstone/repository"
    "dibimbing_golang_capstone/service"
    "dibimbing_golang_capstone/middleware"

    "github.com/gin-gonic/gin"
)

func setupRoutes(r *gin.Engine) {
    userRepository := repository.NewUserRepository()
    userService := service.NewUserService(userRepository)
    userController := controller.NewUserController(userService)

    tripRepository := repository.NewTripRepository()
    tripService := service.NewTripService(tripRepository)
    tripController := controller.NewTripController(tripService)

    bookingRepository := repository.NewBookingRepository()
    bookingService := service.NewBookingService(bookingRepository)
    bookingController := controller.NewBookingController(bookingService)

    authMiddleware := middleware.AuthMiddleware()

    authController := controller.NewAuthController(userRepository)

    r.POST("/register", userController.RegisterUser)
    r.POST("/login", authController.Login)

    auth := r.Group("/auth")
    auth.Use(authMiddleware)
    {
        auth.POST("/trips", tripController.CreateTrip)
        auth.GET("/trips", tripController.GetTripsByCityAndDate)
        auth.POST("/bookings", bookingController.CreateBooking)
        auth.GET("/bookings/:trip_id", bookingController.GetBookingsByTripID)
    }
}