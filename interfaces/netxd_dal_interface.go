package interfaces

import "github.com/sivasenthil09/netxd_dal/netxd_customer_dal/models"

type ICustomer interface {
	CreateCustomer(customer *models.Customer) (*models.DBResponse, error)
}
