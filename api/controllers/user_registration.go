package controllers

import (
	"go-final-project/databases"
	"go-final-project/repositories"
	"go-final-project/structs"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CheckUsernameAvailability(username string) (bool, error) {
	var (
		db    = databases.DbConn
		count int
	)

	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE username = $1", username).Scan(&count)
	if err != nil {
		return false, err
	}

	if count > 0 {
		return false, nil
	}

	return true, nil
}

func CheckEmailAvailability(email string) (bool, error) {
	var (
		db    = databases.DbConn
		count int
	)

	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE email = $1", email).Scan(&count)
	if err != nil {
		return false, err
	}

	if count > 0 {
		return false, nil
	}

	return true, nil
}

func RegisterUser(ctx *gin.Context) {
	var user structs.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		panic(err)
	}

	// Check if username and email are available
	usernameAvailability, err := CheckUsernameAvailability(user.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to check username availability: " + err.Error(),
		})
		return
	}

	emailAvailability, err := CheckEmailAvailability(user.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to check email availability: " + err.Error(),
		})
		return
	}

	if !usernameAvailability || !emailAvailability {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Username or email already taken",
		})
		return
	}

	// Generate password hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to generate password hash: " + err.Error(),
		})
		return
	}

	// Insert hashed password as user password into database
	user.Password = string(hashedPassword)

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	if err := repositories.CreateUser(databases.DbConn, user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}
