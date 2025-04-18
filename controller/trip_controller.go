package controller

import (
	"net/http"
	"dibimbing_golang_capstone/dto"
	"dibimbing_golang_capstone/service"
	"github.com/gin-gonic/gin"
)

type TripController struct {
	tripService service.TripService
}

func NewTripController(tripService service.TripService) TripController {
	return TripController{tripService}
}

func (c *TripController) CreateTrip(ctx *gin.Context) {
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
