package models

import (
	"gorm.io/gorm"
)

type OrderStatus string

const (
	Pending   OrderStatus = "Pending"
	Shipped   OrderStatus = "Shipped"
	Delivered OrderStatus = "Delivered"
	Canceled  OrderStatus = "Canceled"
)

type Order struct {
	gorm.Model
	UserID     uint        `json:"user_id"`
	Status     OrderStatus `json:"status" gorm:"default:'Pending'"`
	OrderItems []OrderItem `json:"order_items"`
}

type OrderItem struct {
	gorm.Model
	OrderID   uint `json:"order_id"`
	ProductID uint `json:"product_id"`
	Quantity  uint `json:"quantity"`
}
