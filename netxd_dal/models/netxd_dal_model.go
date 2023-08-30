package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Customer struct {
	CustomerID primitive.ObjectID `json:"customer" bson:"customer_id"`
	FirstName  string             `json:"first_name" bson:"first_name"`
	LastName   string             `json:"last_name" bson:"last_name"`
	BankID     int64              `json:"bank_id" bson:"bank_id"`
	Balance    float64            `json:"balance" bson:"balance"`
	CreatedAt  string             `json:"created_at" bson:"created_at"`
	UpdateAt   string             `json:"update_at" bson:"update_at"`
	IsActive   bool               `json:"is_active" bson:"is_active"`
}
type DBResponse struct {
	CreatedAt  string             `json:"created_at" bson:"created_at"`
	CustomerID primitive.ObjectID `json:"customer" bson:"customer_id"`
}
