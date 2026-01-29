package order

import "github.com/fahmiabd/go-order-api/internal/model"

type IOrderRepository interface {
	Create(order *model.Order) error
	FindByID(id uint) (*model.Order, error)
	FindByUser(
		userID uint,
		limit int,
		offset int,
	) ([]model.Order, int64, error)
	Update(order *model.Order) error
	Delete(id uint) error
}
