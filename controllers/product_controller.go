package controllers

import (
	"net/http"
	"product-service-app/models"
	"product-service-app/services"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	ProductService *services.ProductService
}

func NewProductController(productService *services.ProductService) *ProductController {
	return &ProductController{ProductService: productService}
}

// CreateProduct godoc
// @Summary Create a new product
// @Description Create a new product with the input payload
// @Tags products
// @Accept json
// @Produce json
// @Param product body models.CreateProductInput true "Product Data"
// @Success 201 {object} models.Product
// @Router /api/products [post]
// @Security BearerAuth
func (pc *ProductController) CreateProduct(c *gin.Context) {
	var input models.CreateProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, err := pc.ProductService.CreateProduct(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create product"})
		return
	}

	c.JSON(http.StatusCreated, product)
}

// ListProducts godoc
// @Summary List all products for a user
// @Description Get all products placed by the authenticated user
// @Tags products
// @Accept json
// @Produce json
// @Success 200 {array} models.Product
// @Router /api/products [get]
// @Security BearerAuth
func (pc *ProductController) GetAllProducts(c *gin.Context) {
	products, err := pc.ProductService.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve products"})
		return
	}

	c.JSON(http.StatusOK, products)
}

// GetProduct godoc
// @Summary Get a product by ID
// @Description Retrieve a single product by its ID
// @Tags products
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} models.Product
// @Router /api/products/{id} [get]
// @Security BearerAuth
func (pc *ProductController) GetProductByID(c *gin.Context) {
	id := c.Param("id")
	product, err := pc.ProductService.GetProductByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}

// UpdateProduct godoc
// @Summary update existing product
// @Description update existing product with the input payload
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param product body models.UpdateProductInput true "Product Data"
// @Success 201 {object} models.Product
// @Router /api/products/:id [put]
// @Security BearerAuth
func (pc *ProductController) UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var input models.UpdateProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, err := pc.ProductService.UpdateProduct(id, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update product"})
		return
	}

	c.JSON(http.StatusOK, product)
}

// DeleteProduct godoc
// @Summary Delete a product by ID
// @Description delete a single product by its ID
// @Tags products
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} models.Product
// @Router /api/products/{id} [delete]
// @Security BearerAuth
func (pc *ProductController) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	err := pc.ProductService.DeleteProduct(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete product"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
