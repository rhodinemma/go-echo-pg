package database

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/rhodinemma/go-echo-pg/internal/dberrors"
	"github.com/rhodinemma/go-echo-pg/internal/models"
	"gorm.io/gorm"
)

func (c Client) GetAllProducts(ctx context.Context, vendorID string) ([]models.Product, error) {
	var products []models.Product
	result := c.DB.WithContext(ctx).Where(models.Product{VendorID: vendorID}).Find(&products)
	return products, result.Error
}

func (c Client) AddProduct(ctx context.Context, product *models.Product) (*models.Product, error) {
	product.ProductID = uuid.NewString()
	result := c.DB.WithContext(ctx).Create(&product)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &dberrors.ConflictError{}
		}
		return nil, result.Error
	}

	return product, nil
}

func (c Client) GetProductById(ctx context.Context, ID string) (*models.Product, error) {
	product := &models.Product{}
	result := c.DB.WithContext(ctx).Where(&models.Product{ProductID: ID}).First(&product)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, &dberrors.NotFoundError{Entity: "product", ID: ID}
		}
		return nil, result.Error
	}

	return product, nil
}
