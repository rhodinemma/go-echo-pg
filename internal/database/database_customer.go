package database

import (
	"context"

	"github.com/rhodinemma/go-echo-pg/internal/models"
)

func (c Client) GetAllCustomers(ctx context.Context, emailAddress string) ([]models.Customer, error) {
	var customers []models.Customer

	result := c.DB.WithContext(ctx).Where(models.Customer{Email: emailAddress}).Find(&customers)
	return customers, result.Error
}
