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

func GetUserEvent(ctx *gin.Context) {
	var (
		result gin.H
	)

	user_event, err := repositories.GetUserEvent(databases.DbConn)
	if err != nil {
		result = gin.H{
			"status":  "error",
			"message": err,
		}
	} else {
		result = gin.H{
			"status": "success",
			"data":   user_event,
		}

	}

	ctx.JSON(http.StatusOK, result)
}

func CreateUserEvent(ctx *gin.Context) {
	var user_event structs.UserEvent

	if err := ctx.ShouldBindJSON(&user_event); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	user_event.RegisDate = time.Now()
	user_event.CreatedAt = time.Now()
	user_event.UpdatedAt = time.Now()

	if err := repositories.CreateUserEvent(databases.DbConn, user_event); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User in event created successfully"})
}

func UpdateUserEvent(ctx *gin.Context) {
	var user_event structs.UserEvent
	id, _ := strconv.Atoi(ctx.Param("id"))

	if err := ctx.ShouldBindJSON(&user_event); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	user_event.UpdatedAt = time.Now()

	if err := repositories.UpdateUserEvent(databases.DbConn, id, user_event); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User in event updated successfully"})
}

func DeleteUserEvent(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	if err := repositories.DeleteUserEvent(databases.DbConn, id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User in event deleted successfully"})
}

func FindUserEvent(ctx *gin.Context) {
	var user structs.UserEvent
	id, _ := strconv.Atoi(ctx.Param("id"))

	user, err := repositories.FindUserEvent(databases.DbConn, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": user})
}
