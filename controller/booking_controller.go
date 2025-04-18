package controller

import (
	"net/http"
	"dibimbing_golang_capstone/dto"
	"dibimbing_golang_capstone/service"
	"github.com/gin-gonic/gin"
)

type BookingController struct {
	bookingService service.BookingService
}

func NewBookingController(bookingService service.BookingService) BookingController {
	return BookingController{bookingService}
}

func (c *BookingController) CreateBooking(ctx *gin.Context) {
	var bookingDTO dto.CreateBookingDTO
	if err := ctx.ShouldBindJSON(&bookingDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := ctx.MustGet("userID").(uint)

	booking, err := c.bookingService.CreateBooking(bookingDTO, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"booking": booking})
}

func (c *BookingController) GetBookingsByTripID(ctx *gin.Context) {
	tripID := ctx.Param("trip_id")

	tripIDUint := uint(0)
	if _, err := fmt.Sscanf(tripID, "%d", &tripIDUint); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid trip ID"})
		return
	}

	bookings, err := c.bookingService.GetBookingsByTripID(tripIDUint)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"bookings": bookings})
}
