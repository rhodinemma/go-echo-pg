package database

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/rhodinemma/go-echo-pg/internal/dberrors"
	"github.com/rhodinemma/go-echo-pg/internal/models"
	"gorm.io/gorm"
)

func (c Client) GetAllCustomers(ctx context.Context, emailAddress string) ([]models.Customer, error) {
	var customers []models.Customer

	result := c.DB.WithContext(ctx).Where(models.Customer{Email: emailAddress}).Find(&customers)
	return customers, result.Error
}

func (c Client) AddCustomer(ctx context.Context, customer *models.Customer) (*models.Customer, error) {
	customer.CustomerID = uuid.NewString()
	result := c.DB.WithContext(ctx).Create(&customer)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &dberrors.ConflictError{}
		}

		return nil, result.Error
	}

	return customer, nil
}

func (c Client) GetCustomerById(ctx context.Context, ID string) (*models.Customer, error) {
	customer := &models.Customer{}
	result := c.DB.WithContext(ctx).Where(&models.Customer{CustomerID: ID}).First(&customer)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, &dberrors.NotFoundError{Entity: "customer", ID: ID}
		}
		return nil, result.Error
	}

	return customer, nil
}
