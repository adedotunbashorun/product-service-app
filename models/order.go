package models

type OrderStatus string

const (
	Pending   OrderStatus = "Pending"
	Shipped   OrderStatus = "Shipped"
	Delivered OrderStatus = "Delivered"
	Canceled  OrderStatus = "Canceled"
)

type Order struct {
	BaseModel
	UserID     uint        `json:"user_id"`
	Status     OrderStatus `json:"status" gorm:"default:'Pending'"`
	OrderItems []OrderItem `json:"order_items"`
}

type OrderItem struct {
	BaseModel
	OrderID   uint `json:"order_id"`
	ProductID uint `json:"product_id"`
	Quantity  uint `json:"quantity"`
}
