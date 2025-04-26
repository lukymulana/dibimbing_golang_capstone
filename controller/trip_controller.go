package controller

import (
	"net/http"
	"dibimbing_golang_capstone/dto"
	"dibimbing_golang_capstone/service"
	"dibimbing_golang_capstone/repository"
	"github.com/gin-gonic/gin"
)

type TripController struct {
	tripService service.TripService
}

func NewTripController(tripService service.TripService) TripController {
	return TripController{tripService}
}

func (c *TripController) CreateTrip(ctx *gin.Context) {
	role := ctx.MustGet("role").(string)
	if role != "guide" {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Only guide can create trip"})
		return
	}

	var tripDTO dto.CreateTripDTO
	if err := ctx.ShouldBindJSON(&tripDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := ctx.MustGet("userID").(uint)

	trip, err := c.tripService.CreateTrip(tripDTO, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"trip": trip})
}

func (c *TripController) UpdateTrip(ctx *gin.Context) {
	role := ctx.MustGet("role").(string)
	if role != "guide" {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Only guide can update trip"})
		return
	}

	id := ctx.Param("id")
	var tripDTO dto.CreateTripDTO
	if err := ctx.ShouldBindJSON(&tripDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID := ctx.MustGet("userID").(uint)
	trip, err := c.tripService.UpdateTrip(id, tripDTO, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"trip": trip})
}

func (c *TripController) DeleteTrip(ctx *gin.Context) {
	role := ctx.MustGet("role").(string)
	if role != "guide" {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Only guide can delete trip"})
		return
	}

	id := ctx.Param("id")
	userID := ctx.MustGet("userID").(uint)
	if err := c.tripService.DeleteTrip(id, userID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Trip deleted successfully"})
}

func (c *TripController) GetMyTrips(ctx *gin.Context) {
	userID := ctx.MustGet("userID").(uint)
	trips, err := c.tripService.GetTripsByUserID(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"trips": trips})
}

func (c *TripController) GetTripsByCityAndDate(ctx *gin.Context) {
	city := ctx.Query("city")
	startDate := ctx.Query("start_date")
	endDate := ctx.Query("end_date")

	trips, err := c.tripService.GetTripsByCityAndDate(city, startDate, endDate)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"trips": trips})
}


func (c *TripController) GetAllTrips(ctx *gin.Context) {
	trips, err := c.tripService.GetAllTrips()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"trips": trips})
}

func (c *TripController) GetTripByID(ctx *gin.Context) {
	id := ctx.Param("id")
	trip, err := c.tripService.GetTripByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Hitung sisa kuota
	bookingService := service.NewBookingService(repository.NewBookingRepository())
	bookings, err := bookingService.GetBookingsByTripID(trip.TripID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	sisaKuota := trip.Capacity - len(bookings)

	ctx.JSON(http.StatusOK, gin.H{
		"trip": trip,
		"sisa_kuota": sisaKuota,
	})
}


