package seeder

import (
	"product-service-app/models"

	"gorm.io/gorm"
)

func SeedOrders(db *gorm.DB) error {
	// Sample orders with items
	orders := []models.Order{
		{
			UserID: 1, // Assuming User with ID 1 exists
			Status: "Pending",
			OrderItems: []models.OrderItem{
				{ProductID: 1, Quantity: 2}, // Assuming Product with ID 1 exists
				{ProductID: 2, Quantity: 1}, // Assuming Product with ID 2 exists
			},
		},
		{
			UserID: 2, // Assuming User with ID 2 exists
			Status: "Completed",
			OrderItems: []models.OrderItem{
				{ProductID: 3, Quantity: 1}, // Assuming Product with ID 3 exists
				{ProductID: 4, Quantity: 3}, // Assuming Product with ID 4 exists
			},
		},
		{
			UserID: 1, // Another order for User with ID 1
			Status: "Shipped",
			OrderItems: []models.OrderItem{
				{ProductID: 2, Quantity: 2}, // Assuming Product with ID 2 exists
				{ProductID: 5, Quantity: 1}, // Assuming Product with ID 5 exists
			},
		},
	}

	// Seed each order
	for _, order := range orders {
		// Check if order with same UserID and Status already exists
		var existingOrder models.Order
		db.Where("user_id = ? AND status = ?", order.UserID, order.Status).First(&existingOrder)
		if existingOrder.ID == 0 { // Order not found
			if err := db.Create(&order).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
