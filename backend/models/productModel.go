package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID    primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name  string             `json:"name" bson:"name"`
	Price int                `json:"price" bson:"price"`
	Stock int                `json:"stock" bson:"stock"`
}
