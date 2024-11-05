package controllers

import (
	"net/http"
	"user-management-mysql/models"
	"user-management-mysql/services"

	"github.com/gin-gonic/gin"
)

type RoleController struct {
	RoleService *services.RoleService
}

func NewRoleController(roleService *services.RoleService) *RoleController {
	return &RoleController{RoleService: roleService}
}

func (ctrl *RoleController) CreateRole(c *gin.Context) {
	var role models.CreateRoleInput
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdRole, err := ctrl.RoleService.CreateRole(role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdRole)
}

func (ctrl *RoleController) GetRoles(c *gin.Context) {
	roles, err := ctrl.RoleService.GetRoles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, roles)
}

func (ctrl *RoleController) UpdateRole(c *gin.Context) {
	id := c.Param("id")
	var role models.UpdateRoleInput
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedRole, err := ctrl.RoleService.UpdateRole(id, role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedRole)
}

func (ctrl *RoleController) DeleteRole(c *gin.Context) {
	roleID := c.Param("id")

	err := ctrl.RoleService.DeleteRole(roleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Role deleted successfully"})
}
