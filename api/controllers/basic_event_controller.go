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

func GetEvent(ctx *gin.Context) {
	var (
		result gin.H
	)

	event, err := repositories.GetEvent(databases.DbConn)
	if err != nil {
		result = gin.H{
			"status":  "error",
			"message": err,
		}
	} else {
		result = gin.H{
			"status": "success",
			"data":   event,
		}

	}

	ctx.JSON(http.StatusOK, result)
}

func CreateEvent(ctx *gin.Context) {
	var event structs.Event

	if err := ctx.ShouldBindJSON(&event); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	event.CreatedAt = time.Now()
	event.UpdatedAt = time.Now()

	if err := repositories.CreateEvent(databases.DbConn, event); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create event"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Event created successfully"})
}

func UpdateEvent(ctx *gin.Context) {
	var event structs.Event
	id, _ := strconv.Atoi(ctx.Param("id"))

	if err := ctx.ShouldBindJSON(&event); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	event.UpdatedAt = time.Now()

	if err := repositories.UpdateEvent(databases.DbConn, id, event); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update event"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "event updated successfully"})
}

func DeleteEvent(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	if err := repositories.DeleteEvent(databases.DbConn, id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete event"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}

func FindEvent(ctx *gin.Context) {
	var event structs.Event
	id, _ := strconv.Atoi(ctx.Param("id"))

	event, err := repositories.FindEvent(databases.DbConn, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get event"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": event})
}
