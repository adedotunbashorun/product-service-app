package routes

import (
	"user-management-mysql/controllers"
	"user-management-mysql/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.RouterGroup, userController *controllers.UserController) {
	auth := router.Group("/auth")
	{
		auth.POST("/register", userController.Register)
		auth.POST("/login", userController.Login)
	}

	user := router.Group("/user")
	user.Use(middleware.AuthMiddleware()) // Apply authentication middleware
	{
		user.GET("/me", userController.GetCurrentUser)
	}
}
