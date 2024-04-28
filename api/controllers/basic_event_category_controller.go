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

func GetEventCategory(ctx *gin.Context) {
	var (
		result gin.H
	)

	event, err := repositories.GetEventCategory(databases.DbConn)
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

func CreateEventCategory(ctx *gin.Context) {
	var event_category structs.Event_Categories

	if err := ctx.ShouldBindJSON(&event_category); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	event_category.CreatedAt = time.Now()
	event_category.UpdatedAt = time.Now()

	if err := repositories.CreateEventCategory(databases.DbConn, event_category); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create event category"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Event category created successfully"})
}

func UpdateEventCategory(ctx *gin.Context) {
	var event_category structs.Event_Categories
	id, _ := strconv.Atoi(ctx.Param("id"))

	if err := ctx.ShouldBindJSON(&event_category); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	event_category.UpdatedAt = time.Now()

	if err := repositories.UpdateEventCategory(databases.DbConn, id, event_category); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update event category"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Event category updated successfully"})
}

func DeleteEventCategory(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	if err := repositories.DeleteEventCategory(databases.DbConn, id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete event category"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Event category deleted successfully"})
}

func FindEventCategory(ctx *gin.Context) {
	var event_category structs.Event_Categories
	id, _ := strconv.Atoi(ctx.Param("id"))

	event_category, err := repositories.FindEventCategory(databases.DbConn, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get event category"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": event_category})
}
