package routes

import (
	"user-management-mysql/controllers"
	"user-management-mysql/middleware"

	"github.com/gin-gonic/gin"
)

// RegisterOrderRoutes registers the order-related routes.
func RegisterOrderRoutes(router *gin.RouterGroup, orderController *controllers.OrderController) {
	order := router.Group("/orders")
	order.Use(middleware.AuthMiddleware()) // Apply authentication middleware
	{
		order.POST("/", orderController.PlaceOrder)
		order.GET("/", orderController.ListOrders)
		order.PUT("/:id/cancel", orderController.CancelOrder)
	}
	order.Use(middleware.AdminRoleMiddleware()) // Apply authentication middleware
	{
		order.PUT("/:id/status", orderController.UpdateOrderStatus)
	}
}
