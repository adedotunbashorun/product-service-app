package repositories

import (
	"product-service-app/models"

	"gorm.io/gorm"
)

// ProductRepository provides access to product data.
type ProductRepository struct {
	DB *gorm.DB
}

// NewProductRepository creates a new ProductRepository.
func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

// Create saves a new product to the database.
func (r *ProductRepository) Create(product *models.Product) error {
	return r.DB.Create(product).Error
}

// GetAll retrieves all products from the database.
func (r *ProductRepository) GetAll() ([]models.Product, error) {
	var products []models.Product
	err := r.DB.Find(&products).Error
	return products, err
}

// GetByID retrieves a product by ID.
func (r *ProductRepository) GetByID(id string) (models.Product, error) {
	var product models.Product
	err := r.DB.First(&product, "id = ?", id).Error // Change ID type to string
	return product, err
}

// Update modifies an existing product.
func (r *ProductRepository) Update(product *models.Product) error {
	return r.DB.Save(product).Error
}

// Delete removes a product from the database.
func (r *ProductRepository) Delete(id string) error {
	return r.DB.Delete(&models.Product{}, "id = ?", id).Error // Change ID type to string
}
