package controllers

import (
	"net/http"
	"user-management-mysql/models"
	"user-management-mysql/services"

	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	service *services.OrderService
}

func NewOrderController(service *services.OrderService) *OrderController {
	return &OrderController{service: service}
}

// PlaceOrder godoc
// @Summary Place a new order
// @Description Place an order for one or more products
// @Tags orders
// @Accept json
// @Produce json
// @Param order body []models.OrderItem true "Order Items"
// @Success 201 {object} models.Order
// @Router /api/orders [post]
// @Security BearerAuth
func (ctrl *OrderController) PlaceOrder(c *gin.Context) {
	var items []models.OrderItem
	if err := c.ShouldBindJSON(&items); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.MustGet("userID").(uint)
	order, err := ctrl.service.PlaceOrder(userID, items)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, order)
}

// ListOrders godoc
// @Summary List all orders for a user
// @Description Get all orders placed by the authenticated user
// @Tags orders
// @Accept json
// @Produce json
// @Success 200 {array} models.Order
// @Router /api/orders [get]
// @Security BearerAuth
func (ctrl *OrderController) ListOrders(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	orders, err := ctrl.service.ListOrders(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, orders)
}

// CancelOrder godoc
// @Summary Cancel a pending order
// @Description Cancel an order if it's still pending
// @Tags orders
// @Param id path int true "Order ID"
// @Success 200 {string} string "Order canceled successfully"
// @Router /api/orders/{id}/cancel [put]
// @Security BearerAuth
func (ctrl *OrderController) CancelOrder(c *gin.Context) {
	orderID, _ := strconv.Atoi(c.Param("id"))
	err := ctrl.service.CancelOrder(uint(orderID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Order canceled successfully"})
}

// UpdateOrderStatus godoc
// @Summary Update the status of an order
// @Description Update order status (admin only)
// @Tags orders
// @Param id path int true "Order ID"
// @Param status body string true "New Status"
// @Success 200 {string} string "Status updated successfully"
// @Router /api/orders/{id}/status [put]
// @Security BearerAuth
func (ctrl *OrderController) UpdateOrderStatus(c *gin.Context) {
	orderID, _ := strconv.Atoi(c.Param("id"))
	var request struct {
		Status models.OrderStatus `json:"status"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := ctrl.service.UpdateOrderStatus(uint(orderID), request.Status)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Status updated successfully"})
}
