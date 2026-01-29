package order

import "github.com/fahmiabd/go-order-api/internal/models"

type IOrderRepository interface {
	Create(order *models.Order) error
	FindByID(id uint) (*models.Order, error)
	FindByUser(
		userID uint,
		limit int,
		offset int,
	) ([]models.Order, int64, error)
	Update(order *models.Order) error
	Delete(id uint) error
}
