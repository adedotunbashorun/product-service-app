package seeder

import (
	"product-service-app/models"

	"gorm.io/gorm"
)

func SeedProducts(db *gorm.DB) error {
	products := []models.Product{
		{Name: "Laptop", Description: "High-performance laptop", Price: 1200.00},
		{Name: "Smartphone", Description: "Latest model smartphone", Price: 800.00},
		{Name: "Headphones", Description: "Noise-canceling headphones", Price: 150.00},
		{Name: "Smartwatch", Description: "Fitness tracking smartwatch", Price: 200.00},
		{Name: "Camera", Description: "Digital camera with 4K video", Price: 500.00},
		{Name: "Tablet", Description: "Portable and powerful tablet", Price: 300.00},
		{Name: "Monitor", Description: "4K UHD monitor", Price: 400.00},
		{Name: "Keyboard", Description: "Mechanical keyboard", Price: 100.00},
		{Name: "Mouse", Description: "Wireless ergonomic mouse", Price: 50.00},
		{Name: "Printer", Description: "All-in-one laser printer", Price: 250.00},
	}

	for _, product := range products {
		var existingProduct models.Product
		// Check if the product already exists
		db.Where("name = ?", product.Name).First(&existingProduct)
		if existingProduct.ID == 0 { // Product not found
			if err := db.Create(&product).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
