package routes

import (
	"product-service-app/controllers"
	"product-service-app/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.RouterGroup, userController *controllers.UserController) {
	auth := router.Group("/auth")
	{
		auth.POST("/register", userController.Register)
		auth.POST("/login", userController.Login)
	}

	users := router.Group("/users")
	users.Use(middleware.AuthMiddleware()) // Apply authentication middleware
	{
		users.GET("/me", userController.GetCurrentUser)
	}
}
