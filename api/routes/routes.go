package routes

import (
	"go-final-project/api/controllers"
	"go-final-project/api/middlewares"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	r := gin.Default()

	// Public Routes
	r.POST("/register-user", controllers.RegisterUser)
	r.GET("/login-user", middlewares.BasicAuthMiddleware, controllers.ValidateUser)
	r.POST("/register-event/:id", controllers.RegisterEvent)
	r.GET("/registered-events/:id", middlewares.JWTAuthMiddleware, controllers.GetRegisteredEvents)

	// Basic User Routes
	r.GET("/user", controllers.GetUser)
	r.POST("/user", controllers.CreateUser)
	r.PUT("/update-user/:id", controllers.UpdateUser)
	r.DELETE("/delete-user/:id", controllers.DeleteUser)

	// Basic Event Category Routes
	r.GET("/event-category", controllers.GetEventCategory)
	r.POST("/event-category", controllers.CreateEventCategory)
	r.PUT("/update-event-category/:id", controllers.UpdateEventCategory)
	r.DELETE("/delete-event-category/:id", controllers.DeleteEventCategory)

	// Basic Event Routes
	r.GET("/event", controllers.GetEvent)
	r.POST("/event", controllers.CreateEvent)
	r.PUT("/update-event/:id", controllers.UpdateEvent)
	r.DELETE("/delete-event/:id", controllers.DeleteEvent)

	// Basic User Event Routes
	r.GET("/user-event", controllers.GetUserEvent)
	r.POST("/user-event", controllers.CreateUserEvent)
	r.PUT("/update-user-event/:id", controllers.UpdateUserEvent)
	r.DELETE("/delete-user-event/:id", controllers.DeleteUserEvent)

	return r
}
