package middlewares

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuthMiddleware(ctx *gin.Context) {
	// Get token from cookie
	tokenString, err := ctx.Cookie("token")
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Verify token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret"), nil // Replace "secret" with your actual secret key
	})
	if err != nil || !token.Valid {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Extract claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		// Check if token is expired
		exp, ok := claims["exp"].(float64)
		if !ok || time.Now().Unix() > int64(exp) {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Set user in context
		ctx.Set("user", claims["username"])

		// Continue
		ctx.Next()

		// Print claims
		fmt.Println(claims["foo"], claims["nbf"])
	} else {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}
}
