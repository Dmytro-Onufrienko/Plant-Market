package orders

import "time"

type Order struct {
	ID              int             `json:"id"`
	UserID          string          `json:"user_id"`
	Status          string          `json:"status"`
	TotalPrice      int             `json:"total_price"` // в центах
	ShippingAddress ShippingAddress `json:"shipping_address"`
	CreatedAt       time.Time       `json:"created_at"`
}

type OrderWithItems struct {
	Order
	Items []OrderItem `json:"items"`
}

type OrderWithUser struct {
	Order
	UserEmail string `json:"user_email"`
	UserName  string `json:"user_name"`
}

type ShippingAddress struct {
	FullName string `json:"full_name"`
	Street   string `json:"street"`
	City     string `json:"city"`
	Zip      string `json:"zip"`
	Country  string `json:"country"`
}

type OrderItem struct {
	ID              int `json:"id"`
	OrderID         int `json:"order_id"`
	PlantID         int `json:"plant_id"`
	Quantity        int `json:"quantity"`
	PriceAtPurchase int `json:"price_at_purchase"` // в центах
}

type OrderItemWithPlant struct {
	OrderItem
	PlantName   string   `json:"plant_name"`
	PlantSlug   string   `json:"plant_slug"`
	PlantImages []string `json:"plant_images"`
}

type OrderWithItemsDetailed struct {
	Order
	Items []OrderItemWithPlant `json:"items"`
}

type CheckoutRequest struct {
	Items    []ItemInput     `json:"items"    binding:"required,min=1,dive"`
	Shipping ShippingAddress `json:"shipping" binding:"required"`
}

type ItemInput struct {
	PlantID  int `json:"plant_id"  binding:"required"`
	Quantity int `json:"quantity"  binding:"required,min=1"`
}

type UpdateStatusRequest struct {
	Status string `json:"status" binding:"required,oneof=pending paid shipped delivered"`
}
