package product

import "github.com/fahmiabd/go-order-api/internal/models"

type IProductRepository interface {
	FindByID(id uint) (*models.Product, error)
}
