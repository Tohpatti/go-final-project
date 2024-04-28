package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ValidateUser(ctx *gin.Context) {
	user, _ := ctx.Get("user")

	ctx.JSON(http.StatusOK, gin.H{
		"message": user.(string) + " is logged in.",
	})
}
