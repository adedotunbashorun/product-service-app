package repositories

import (
	"product-service-app/models"

	"gorm.io/gorm"
)

// ProductRepository provides access to product data.
type ProductRepository struct {
	BaseRepository[models.Product]
}

// NewProductRepository creates a new ProductRepository.
func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		BaseRepository: BaseRepository[models.Product]{DB: db},
	}
}

// GetByID retrieves a product by ID.
func (r *ProductRepository) GetByID(id string) (models.Product, error) {
	var product models.Product
	err := r.DB.First(&product, "id = ?", id).Error // Change ID type to string
	return product, err
}
