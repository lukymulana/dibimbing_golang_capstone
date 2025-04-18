package dto

type CreateTripDTO struct {
	City        string  `json:"city" binding:"required"`
	StartDate   string  `json:"start_date" binding:"required"`
	EndDate     string  `json:"end_date" binding:"required"`
	Capacity    int     `json:"capacity" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
	Description string  `json:"description"`
}
