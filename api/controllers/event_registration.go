package controllers

import (
	"go-final-project/databases"
	"go-final-project/repositories"
	"go-final-project/structs"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func RegisterEvent(ctx *gin.Context) {
	var user_event structs.UserEvent
	if err := ctx.ShouldBindJSON(&user_event); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input",
		})
		return
	}

	user_event.RegisDate = time.Now()
	user_event.CreatedAt = time.Now()
	user_event.UpdatedAt = time.Now()

	err := repositories.CreateUserEvent(databases.DbConn, user_event)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to register user to event: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User registered to event successfully",
	})
}
