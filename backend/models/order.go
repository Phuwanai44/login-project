package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ShippingAddress struct {
	FullName string `json:"fullName" bson:"fullName"`
	Phone    string `json:"phone" bson:"phone"`
	Address  string `json:"address" bson:"address"`
	City     string `json:"city" bson:"city"`
}

type OrderItem struct {
	ProductID string `json:"productId" bson:"productId"`
	Name      string `json:"name" bson:"name"`
	Color     string `json:"color" bson:"color"`
	Size      string `json:"size" bson:"size"`
	Quantity  int    `json:"quantity" bson:"quantity"`
	Price     int    `json:"price" bson:"price"`
}

type Order struct {
	ID              primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID          string             `json:"userId" bson:"userId"`
	Items           []OrderItem        `json:"items" bson:"items"`
	TotalPrice      int                `json:"totalPrice" bson:"totalPrice"`
	ShippingAddress ShippingAddress    `json:"shippingAddress" bson:"shippingAddress"`
	Status          string             `json:"status" bson:"status"` // pending, shipped, delivered, cancelled
	CreatedAt       time.Time          `json:"createdAt" bson:"createdAt"`
}

type CreateOrderInput struct {
	Items           []OrderItem     `json:"items"`
	ShippingAddress ShippingAddress `json:"shippingAddress"`
}
