package controllers

import (
	"net/http"
	"product-service-app/models"
	"product-service-app/services"

	"github.com/gin-gonic/gin"
)

type RoleController struct {
	RoleService *services.RoleService
}

func NewRoleController(roleService *services.RoleService) *RoleController {
	return &RoleController{RoleService: roleService}
}

// CreateRole godoc
// @Summary Create a new Role
// @Description Create a new Role with the input payload
// @Tags Roles
// @Accept json
// @Produce json
// @Param Role body models.CreateRoleInput true "Role Data"
// @Success 201 {object} models.Role
// @Router /api/roles [post]
// @Security BearerAuth
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

// ListRoles godoc
// @Summary List all Roles for a user
// @Description Get all Roles placed by the authenticated user
// @Tags Roles
// @Accept json
// @Produce json
// @Success 200 {array} models.Role
// @Router /api/roles [get]
// @Security BearerAuth
func (ctrl *RoleController) GetRoles(c *gin.Context) {
	roles, err := ctrl.RoleService.GetRoles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, roles)
}

// UpdateRole godoc
// @Summary update existing Role
// @Description update existing Role with the input payload
// @Tags Roles
// @Accept json
// @Produce json
// @Param id path int true "Role ID"
// @Param Role body models.UpdateRoleInput true "Role Data"
// @Success 201 {object} models.Role
// @Router /api/roles/:id [put]
// @Security BearerAuth
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

// DeleteRole godoc
// @Summary Delete a Role by ID
// @Description delete a single Role by its ID
// @Tags Roles
// @Produce json
// @Param id path int true "Role ID"
// @Success 200 {object} models.Role
// @Router /api/roles/{id} [delete]
// @Security BearerAuth
func (ctrl *RoleController) DeleteRole(c *gin.Context) {
	roleID := c.Param("id")

	err := ctrl.RoleService.DeleteRole(roleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Role deleted successfully"})
}
