package interfaces

import (
	"netxd_dal/models"
)

type Icustomer interface {
	CreateCustomer(*models.Customer) (*models.DBResponse,error)
}
