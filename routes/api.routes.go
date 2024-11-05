package routes

import (
	"product-service-app/controllers"

	"github.com/gin-gonic/gin"
)

func SetupAPIRoutes(router *gin.Engine,
	userController *controllers.UserController,
	roleController *controllers.RoleController,
	productController *controllers.ProductController,
	orderController *controllers.OrderController,
) {
	api := router.Group("/api")
	{
		RegisterUserRoutes(api, userController)
		RegisterRoleRoutes(api, roleController)
		RegisterProductRoutes(api, productController)
		RegisterOrderRoutes(api, orderController)
	}
}
