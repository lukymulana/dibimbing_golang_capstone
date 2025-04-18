package entity

type Booking struct {
	BookingID     uint   `gorm:"primaryKey"`
	UserID        uint   `gorm:"foreignKey:UserID"`
	TripID        uint   `gorm:"foreignKey:TripID"`
	BookingStatus string `gorm:"enum:'waiting,done,cancel'"`
	CreatedAt     string
}
