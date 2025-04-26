package middleware

import (
    "fmt"
    "net/http"
    "strings"
    "time"

    // "dibimbing_golang_capstone/config"
    "dibimbing_golang_capstone/entity"
    // "dibimbing_golang_capstone/repository"

    "github.com/golang-jwt/jwt/v5"
    "github.com/gin-gonic/gin"
)

var jwtSecret = []byte("12345")

func AuthMiddleware() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        authHeader := ctx.GetHeader("Authorization")
        if authHeader == "" {
            ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
            ctx.Abort()
            return
        }

        tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
            }
            return jwtSecret, nil
        })

        if err != nil {
            ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            ctx.Abort()
            return
        }

        if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
            userID := uint(claims["user_id"].(float64))
            role, ok := claims["role"].(string)
            if !ok {
                ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Role not found in token"})
                ctx.Abort()
                return
            }
            ctx.Set("userID", userID)
            ctx.Set("role", role)
            ctx.Next()
        } else {
            ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
            ctx.Abort()
        }
    }
}

func GenerateToken(user *entity.User) (string, error) {
    claims := jwt.MapClaims{
        "user_id": user.UserID,
        "role":    user.Role,
        "exp":     time.Now().Add(time.Hour * 72).Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtSecret)
}