package dto

import "time"

type CreateBookingDTO struct {
	TripID uint `json:"trip_id" binding:"required"`
}

// Tambahkan DTO lain di bawah ini jika perlu

type UpdateBookingStatusDTO struct {
    Status string `json:"status" binding:"required"`
}


type BookingDetailDTO struct {
    BookingID     uint   `json:"booking_id"`
    BookingStatus string `json:"booking_status"`
    CreatedAt     time.Time `json:"created_at"`
    // Info User
    UserID        uint   `json:"user_id"`
    Username      string `json:"username"`
    Email         string `json:"email"`
    // Info Trip
    TripID        uint   `json:"trip_id"`
    City          string `json:"city"`
    StartDate     string `json:"start_date"`
    EndDate       string `json:"end_date"`
}

