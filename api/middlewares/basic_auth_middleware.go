package middlewares

import (
	"go-final-project/databases"
	"go-final-project/repositories"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func BasicAuthMiddleware(ctx *gin.Context) {
	// Retrieve basic auth credentials
	username, password, ok := ctx.Request.BasicAuth()
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid credentials",
		})
		return
	}

	// Get user by username from database
	user, err := repositories.FindUser(databases.DbConn, username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get user: " + err.Error(),
		})
		return
	}

	// Compare sent in password with stored hash password
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid credentials",
		})
		return
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to generate token: " + err.Error(),
		})
		return
	}

	// Set token as cookie
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("token", tokenString, int(time.Hour*24), "/", "localhost", false, true)

	ctx.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})

	// ctx.Next()
}
