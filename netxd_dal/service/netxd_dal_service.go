package service

import (
	"context"
	"netxd_dal/interfaces"
	"netxd_dal/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CustomerService struct {
	CustomerCollection *mongo.Collection
	ctx                context.Context
}

func InitCustomerService(collection *mongo.Collection, ctx context.Context) interfaces.Icustomer {
	return &CustomerService{collection, ctx}
}
func (p *CustomerService) CreateCustomer(Customer *models.Customer) (*models.DBResponse, error) {
	// Customer.CustomerID=123
	// Customer.FirstName="siva"
	// Customer.LastName="senthil"
	// Customer.BankID=1234
	// Customer.CreatedAt="29.08.2023"
	// Customer.UpdateAt="29.08.2023"
	// Customer.IsActive=true
	res, err := p.CustomerCollection.InsertOne(p.ctx, &Customer)
	if err != nil {
		return  nil,err
	}
	Response := &models.DBResponse{
		CustomerID: res.InsertedID.(primitive.ObjectID),
		CreatedAt: Customer.CreatedAt,
	}
	return Response,nil
}
