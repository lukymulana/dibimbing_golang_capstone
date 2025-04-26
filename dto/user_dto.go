package dto

type CreateUserDTO struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Role     string `json:"role" binding:"required,oneof=guide traveler"`
}
type LoginDTO struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}