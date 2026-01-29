package order

import "github.com/fahmiabd/go-order-api/internal/models"

type IOrderService interface {
	Create(userID, productID uint, quantity int) (*models.Order, error)
	GetByUser(userID uint, limit int, offset int) ([]models.Order, int64, error)
}
