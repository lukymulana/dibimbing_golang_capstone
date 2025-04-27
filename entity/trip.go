package entity

import "time"
type Trip struct {
	TripID      uint   `gorm:"primaryKey"`
	UserID      uint   `gorm:"foreignKey:UserID"`
	City        string
	StartDate   string
	EndDate     string
	Capacity    int
	Price       float64
	Description string
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`
	DeletedAt   *time.Time     `gorm:"index"`
}
