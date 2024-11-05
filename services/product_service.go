package services

import (
	"user-management-mysql/models"
	"user-management-mysql/repositories"
)

// ProductService provides methods to manage products.
type ProductService struct {
	productRepo *repositories.ProductRepository
}

// NewProductService creates a new ProductService.
func NewProductService(productRepo *repositories.ProductRepository) *ProductService {
	return &ProductService{productRepo: productRepo}
}

// CreateProduct creates a new product.
func (s *ProductService) CreateProduct(input models.CreateProductInput) (models.Product, error) {
	product := models.Product{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
	}
	err := s.productRepo.Create(&product)
	return product, err
}

// GetAllProducts retrieves all products.
func (s *ProductService) GetAllProducts() ([]models.Product, error) {
	return s.productRepo.GetAll()
}

// GetProductByID retrieves a product by its ID.
func (s *ProductService) GetProductByID(id string) (models.Product, error) {
	return s.productRepo.GetByID(id)
}

// UpdateProduct updates an existing product.
func (s *ProductService) UpdateProduct(id string, input models.UpdateProductInput) (models.Product, error) {
	product, err := s.productRepo.GetByID(id)
	if err != nil {
		return product, err
	}

	if input.Name != nil {
		product.Name = *input.Name
	}
	if input.Description != nil {
		product.Description = *input.Description
	}
	if input.Price != nil {
		product.Price = *input.Price
	}

	err = s.productRepo.Update(&product)
	return product, err
}

// DeleteProduct deletes a product by its ID.
func (s *ProductService) DeleteProduct(id string) error {
	return s.productRepo.Delete(id)
}
