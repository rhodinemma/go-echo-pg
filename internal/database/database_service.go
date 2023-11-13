package database

import (
	"context"

	"github.com/rhodinemma/go-echo-pg/internal/models"
)

func (c Client) GetAllServices(ctx context.Context) ([]models.Service, error) {
	var services []models.Service
	result := c.DB.WithContext(ctx).Find(&services)
	return services, result.Error
}
