package routes

import (
	"user-management-mysql/controllers"

	"github.com/gin-gonic/gin"
)

func SetupAPIRoutes(router *gin.Engine, userController *controllers.UserController, roleController *controllers.RoleController, productController *controllers.ProductController) {
	api := router.Group("/api")
	{
		RegisterUserRoutes(api, userController)
		RegisterRoleRoutes(api, roleController)
		RegisterProductRoutes(api, productController)
	}
}
