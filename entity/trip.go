package entity

type Trip struct {
	TripID      uint   `gorm:"primaryKey"`
	UserID      uint   `gorm:"foreignKey:UserID"`
	City        string
	StartDate   string
	EndDate     string
	Capacity    int
	Price       float64
	Description string
	CreatedAt   string
	UpdatedAt   string
	DeletedAt   string
}
