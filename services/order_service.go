package services

import (
	"errors"
	"product-service-app/models"
	"product-service-app/repositories"
)

type OrderService struct {
	orderRepo *repositories.OrderRepository
}

func NewOrderService(orderRepo *repositories.OrderRepository) *OrderService {
	return &OrderService{orderRepo: orderRepo}
}

func (s *OrderService) PlaceOrder(userID uint, items []models.OrderItem) (*models.Order, error) {
	order := &models.Order{
		UserID:     userID,
		Status:     models.Pending,
		OrderItems: items,
	}
	err := s.orderRepo.CreateOrder(order)
	return order, err
}

func (s *OrderService) ListOrders(userID uint) ([]models.Order, error) {
	return s.orderRepo.FindOrdersByUserID(userID)
}

func (s *OrderService) CancelOrder(orderID uint) error {
	order, err := s.orderRepo.FindOrderByID(orderID)
	if err != nil {
		return err
	}
	if order.Status != models.Pending {
		return errors.New("only pending orders can be canceled")
	}
	return s.orderRepo.UpdateOrderStatus(orderID, models.Canceled)
}

func (s *OrderService) UpdateOrderStatus(orderID uint, status models.OrderStatus) error {
	return s.orderRepo.UpdateOrderStatus(orderID, status)
}
