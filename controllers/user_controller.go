package controllers

import (
	"net/http"
	"product-service-app/models"
	"product-service-app/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{UserService: userService}
}

type registerUserDto struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	RoleID   uint   `json:"role_id" binding: "required"`
}

type loginUserDto struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// RegisterUser godoc
// @Summary Create a new user
// @Description Create a new user with the input payload
// @Tags users
// @Accept json
// @Produce json
// @Param user body registerUserDto true "User Data"
// @Success 201 {object} models.User
// @Router /api/auth/register [post]
func (ctrl *UserController) Register(c *gin.Context) {
	var req registerUserDto
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := ctrl.UserService.Register(req.Username, req.Email, req.Password, req.RoleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// LoginUser godoc
// @Summary login a user
// @Description authenticated user credentials
// @Tags users
// @Accept json
// @Produce json
// @Param user body loginUserDto true "Login Data"
// @Success 201 {object} models.User
// @Router /api/auth/login [post]
func (ctrl *UserController) Login(c *gin.Context) {
	var req loginUserDto

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := ctrl.UserService.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// GetUser godoc
// @Summary Get current logged in user
// @Description Get currently looged in user
// @Tags user
// @Accept  json
// @Produce  json
// @Success 200 {object} models.User
// @Router /api/users/me [get]
func (ctrl *UserController) GetCurrentUser(c *gin.Context) {
	userID, exists := c.MustGet("userID").(uint)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	var user models.User
	err := ctrl.UserService.UserRepo.DB.Preload("Role").First(&user, userID).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}
