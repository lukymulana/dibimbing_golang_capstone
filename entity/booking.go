package entity

import "time"

type Booking struct {
	BookingID     uint   `gorm:"primaryKey"`
	UserID        uint   `gorm:"foreignKey:UserID"`
	TripID        uint   `gorm:"foreignKey:TripID"`
	BookingStatus string `gorm:"enum:'waiting,done,cancel'"`
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	UpdatedAt     time.Time  `gorm:"autoUpdateTime"`
	DeletedAt     *time.Time `gorm:"index"`
}
