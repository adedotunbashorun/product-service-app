package routes

import (
	"product-service-app/controllers"
	"product-service-app/middleware"

	"github.com/gin-gonic/gin"
)

// RegisterProductRoutes registers the product-related routes.
func RegisterProductRoutes(router *gin.RouterGroup, productController *controllers.ProductController) {
	product := router.Group("/products")
	product.Use(middleware.AdminRoleMiddleware()) // Apply authentication middleware
	{
		product.POST("/", productController.CreateProduct)
		product.GET("/", productController.GetAllProducts)
		product.GET("/:id", productController.GetProductByID)
		product.PUT("/:id", productController.UpdateProduct)
		product.DELETE("/:id", productController.DeleteProduct)
	}
}
