package order

import (
	"errors"

	"github.com/fahmiabd/go-order-api/internal/model"
	"github.com/fahmiabd/go-order-api/internal/repositories/order"
)

type orderService struct {
	orderRepo order.IOrderRepository
}

func NewOrderService(orderRepo order.IOrderRepository) IOrderService {
	return &orderService{
		orderRepo: orderRepo,
	}
}

func (s *orderService) Create(
	userID uint,
	amount float64,
) (*model.Order, error) {

	if userID == 0 {
		return nil, errors.New("invalid user")
	}

	if amount <= 0 {
		return nil, errors.New("amount must be greater than zero")
	}

	order := &model.Order{
		UserID: userID,
		Amount: amount,
		Status: "pending",
	}

	if err := s.orderRepo.Create(order); err != nil {
		return nil, err
	}

	return order, nil
}

func (s *orderService) GetByID(
	userID uint,
	orderID uint,
) (*model.Order, error) {

	if userID == 0 || orderID == 0 {
		return nil, errors.New("invalid id")
	}

	order, err := s.orderRepo.FindByID(orderID)
	if err != nil {
		return nil, err
	}

	if order.UserID != userID {
		return nil, errors.New("unauthorized access to order")
	}

	return order, nil
}

func (s *orderService) ListByUser(
	userID uint,
	page int,
	limit int,
) ([]model.Order, int64, error) {

	if userID == 0 {
		return nil, 0, errors.New("invalid user")
	}

	if limit <= 0 {
		limit = 10
	}

	if page <= 0 {
		page = 1
	}

	offset := (page - 1) * limit

	return s.orderRepo.FindByUser(userID, limit, offset)
}

func (s *orderService) UpdateStatus(
	userID uint,
	orderID uint,
	status string,
) (*model.Order, error) {

	validStatus := map[string]bool{
		"pending":   true,
		"paid":      true,
		"cancelled": true,
	}

	if !validStatus[status] {
		return nil, errors.New("invalid order status")
	}

	order, err := s.GetByID(userID, orderID)
	if err != nil {
		return nil, err
	}

	order.Status = status

	if err := s.orderRepo.Update(order); err != nil {
		return nil, err
	}

	return order, nil
}

func (s *orderService) Delete(
	userID uint,
	orderID uint,
) error {

	order, err := s.GetByID(userID, orderID)
	if err != nil {
		return err
	}

	return s.orderRepo.Delete(order.ID)
}
