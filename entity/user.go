package entity

type User struct {
	UserID   uint   `gorm:"primaryKey"`
	Username string `gorm:"unique"`
	Password string
	Email    string `gorm:"unique"`
	Role     string `gorm:"enum:'guide,traveler'"`
}
