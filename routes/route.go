package routes

import (
	"github.com/gin-gonic/gin"
	"dibimbing_golang_capstone/controller"
	"dibimbing_golang_capstone/middleware"
	"dibimbing_golang_capstone/repository"
	"dibimbing_golang_capstone/service"
	"dibimbing_golang_capstone/config"

)

func SetupRoutes(r *gin.Engine) {
	// Initialize repositories
	userRepo := repository.NewUserRepository(config.DB) // penerapan dependencies injection
	tripRepo := repository.NewTripRepository()
	bookingRepo := repository.NewBookingRepository()

	// Initialize services
	userService := service.NewUserService(userRepo)
	tripService := service.NewTripService(tripRepo)
	bookingService := service.NewBookingService(bookingRepo)

	// Initialize controllers
	userController := controller.NewUserController(userService)
	tripController := controller.NewTripController(tripService)
	bookingController := controller.NewBookingController(bookingService)
	authController := controller.NewAuthController(userRepo)

	// Auth middleware
	authMiddleware := middleware.AuthMiddleware()

	// Public routes
	r.POST("/register", userController.RegisterUser)
	r.POST("/login", authController.Login)
	r.GET("/trips", tripController.GetAllTrips)
	r.GET("/trips/:id", tripController.GetTripByID)
	r.GET("/trips/filter", tripController.GetTripsByCityAndDate)

	// Authenticated routes
	auth := r.Group("/auth")
	auth.Use(authMiddleware)
	{
		auth.POST("/trips", tripController.CreateTrip)
		auth.PUT("/trips/:id", tripController.UpdateTrip)
		auth.DELETE("/trips/:id", tripController.DeleteTrip)
		auth.GET("/my-trips", tripController.GetMyTrips)
		auth.POST("/bookings", bookingController.CreateBooking)
		auth.GET("/bookings/:trip_id", bookingController.GetBookingsByTripID)
	}
}