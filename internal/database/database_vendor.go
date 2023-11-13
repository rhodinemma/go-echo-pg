package database

import (
	"context"

	"github.com/rhodinemma/go-echo-pg/internal/models"
)

func (c Client) GetAllVendors(ctx context.Context) ([]models.Vendor, error) {
	var vendors []models.Vendor
	result := c.DB.WithContext(ctx).Find(&vendors)
	return vendors, result.Error
}
