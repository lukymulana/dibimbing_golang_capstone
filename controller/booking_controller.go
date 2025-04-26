package controller

import (
	// "fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"dibimbing_golang_capstone/dto"
	"dibimbing_golang_capstone/service"
)

type BookingController struct {
	bookingService service.BookingService
}

func NewBookingController(bookingService service.BookingService) BookingController {
	return BookingController{bookingService}
}

func (c *BookingController) CreateBooking(ctx *gin.Context) {
	role := ctx.MustGet("role").(string)
	if role != "traveler" {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Only traveler can book a trip"})
		return
	}

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

func (c *BookingController) GetBookingsByGuideID(ctx *gin.Context) {
	role := ctx.MustGet("role").(string)
	if role != "guide" {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Only guide can access this endpoint"})
		return
	}
	guideID := ctx.MustGet("userID").(uint)
	bookings, err := c.bookingService.GetBookingsByGuideID(guideID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, bookings)
}

func (c *BookingController) UpdateBookingStatus(ctx *gin.Context) {
	role := ctx.MustGet("role").(string)
	if role != "guide" {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Only guide can update booking status"})
		return
	}
	guideID := ctx.MustGet("userID").(uint)
	bookingIDStr := ctx.Param("booking_id")
	bookingID, err := strconv.ParseUint(bookingIDStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid booking ID"})
		return
	}
	var updateDTO dto.UpdateBookingStatusDTO
	if err := ctx.ShouldBindJSON(&updateDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if updateDTO.Status != "success" && updateDTO.Status != "waiting" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status value"})
		return
	}
	err = c.bookingService.UpdateBookingStatus(uint(bookingID), guideID, updateDTO.Status)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Booking status updated successfully"})
}

func (c *BookingController) GetBookingsByTripID(ctx *gin.Context) {
	tripID := ctx.Param("trip_id")

	tripIDUint, err := strconv.ParseUint(tripID, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid trip ID"})
		return
	}

	details, err := c.bookingService.GetBookingsByTripID(uint(tripIDUint))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, details)
}
