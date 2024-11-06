package services

import (
	"product-service-app/models"
	"product-service-app/repositories"
)

// ProductService provides methods to manage products.
type ProductService struct {
	BaseService[models.Product]
	productRepository *repositories.ProductRepository
}

// NewProductService creates a new ProductService.
func NewProductService(productRepo *repositories.ProductRepository) *ProductService {
	return &ProductService{
		BaseService:       BaseService[models.Product]{Repository: &productRepo.BaseRepository},
		productRepository: productRepo,
	}
}

// CreateProduct creates a new product.
func (s *ProductService) CreateProduct(input models.CreateProductInput) (models.Product, error) {
	product := models.Product{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
	}
	err := s.productRepository.Create(&product)
	return product, err
}

// UpdateProduct updates an existing product.
func (s *ProductService) UpdateProduct(id string, input models.UpdateProductInput) (models.Product, error) {
	product, err := s.productRepository.GetByID(id)
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

	err = s.productRepository.Update(&product)
	return product, err
}
