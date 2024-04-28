package controllers

import (
	"go-final-project/databases"
	"go-final-project/repositories"
	"go-final-project/structs"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetUser(ctx *gin.Context) {
	var (
		result gin.H
	)

	user, err := repositories.GetUser(databases.DbConn)
	if err != nil {
		result = gin.H{
			"status":  "error",
			"message": err,
		}
	} else {
		result = gin.H{
			"status": "success",
			"data":   user,
		}

	}

	ctx.JSON(http.StatusOK, result)
}

func CreateUser(ctx *gin.Context) {
	var user structs.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	if err := repositories.CreateUser(databases.DbConn, user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

func UpdateUser(ctx *gin.Context) {
	var user structs.User
	id, _ := strconv.Atoi(ctx.Param("id"))

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	user.UpdatedAt = time.Now()

	if err := repositories.UpdateUser(databases.DbConn, id, user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func DeleteUser(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	if err := repositories.DeleteUser(databases.DbConn, id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func FindUser(ctx *gin.Context) {
	var user structs.User
	id := ctx.Param("id")

	user, err := repositories.FindUser(databases.DbConn, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": user})
}
