package controller

import (
	"net/http"
	"dibimbing_golang_capstone/dto"
	"dibimbing_golang_capstone/entity"
	"dibimbing_golang_capstone/repository"
	"dibimbing_golang_capstone/middleware"
	"golang.org/x/crypto/bcrypt"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	userRepository repository.UserRepository
}

func NewAuthController(userRepository repository.UserRepository) AuthController {
	return AuthController{userRepository}
}

func (c *AuthController) Login(ctx *gin.Context) {
	var loginDTO dto.CreateUserDTO
	if err := ctx.ShouldBindJSON(&loginDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := c.userRepository.GetUserByUsername(loginDTO.Username)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginDTO.Password)); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	token, err := middleware.GenerateToken(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
