package repositories

import (
	"product-service-app/models"

	"gorm.io/gorm"
)

type OrderRepository struct {
	Db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{Db: db}
}

func (repo *OrderRepository) CreateOrder(order *models.Order) error {
	return repo.Db.Create(order).Error
}

func (repo *OrderRepository) FindOrdersByUserID(userID uint) ([]models.Order, error) {
	var orders []models.Order
	err := repo.Db.Preload("OrderItems").Where("user_id = ?", userID).Find(&orders).Error
	return orders, err
}

func (repo *OrderRepository) FindOrderByID(orderID uint) (*models.Order, error) {
	var order models.Order
	err := repo.Db.Preload("OrderItems").First(&order, orderID).Error
	return &order, err
}

func (repo *OrderRepository) UpdateOrderStatus(orderID uint, status models.OrderStatus) error {
	return repo.Db.Model(&models.Order{}).Where("id = ?", orderID).Update("status", status).Error
}
