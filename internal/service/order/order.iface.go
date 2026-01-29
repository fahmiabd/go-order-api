package order

import "github.com/fahmiabd/go-order-api/internal/model"

type IOrderService interface {
	Create(userID uint, amount float64) (*model.Order, error)
	GetByID(userID, orderID uint) (*model.Order, error)
	ListByUser(
		userID uint,
		page int,
		limit int,
	) ([]model.Order, int64, error)
	UpdateStatus(
		userID uint,
		orderID uint,
		status string,
	) (*model.Order, error)
	Delete(userID, orderID uint) error
}
