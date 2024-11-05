package routes

import (
	"user-management-mysql/controllers"
	"user-management-mysql/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoleRoutes(router *gin.RouterGroup, roleController *controllers.RoleController) {
	role := router.Group("/roles")
	role.Use(middleware.AdminRoleMiddleware()) // Apply authentication middleware
	{
		role.POST("/", roleController.CreateRole)
		role.GET("/", roleController.GetRoles)
		role.PUT("/:id", roleController.UpdateRole)
		role.DELETE("/:id", roleController.DeleteRole)
	}
}
