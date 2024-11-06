package repositories

import (
	"product-service-app/models"

	"gorm.io/gorm"
)

type OrderRepository struct {
	BaseRepository[models.Order]
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{
		BaseRepository: BaseRepository[models.Order]{DB: db},
	}
}

func (repo *OrderRepository) FindOrdersByUserID(userID uint) ([]models.Order, error) {
	var orders []models.Order
	err := repo.BaseRepository.DB.Preload("OrderItems").Where("user_id = ?", userID).Find(&orders).Error
	return orders, err
}

func (repo *OrderRepository) FindOrderByID(orderID uint) (*models.Order, error) {
	var order models.Order
	err := repo.BaseRepository.DB.Preload("OrderItems").First(&order, orderID).Error
	return &order, err
}

func (repo *OrderRepository) UpdateOrderStatus(orderID uint, status models.OrderStatus) error {
	return repo.BaseRepository.DB.Model(&models.Order{}).Where("id = ?", orderID).Update("status", status).Error
}
