package order

import (
	"errors"

	"github.com/fahmiabd/go-order-api/internal/models"
	"github.com/fahmiabd/go-order-api/internal/repositories/order"
	"github.com/fahmiabd/go-order-api/internal/repositories/product"
)

type orderService struct {
	orderRepo   order.IOrderRepository
	productRepo product.IProductRepository
}

func NewOrderService(orderRepo order.IOrderRepository, productRepo product.IProductRepository) IOrderService {
	return &orderService{
		orderRepo:   orderRepo,
		productRepo: productRepo,
	}
}

func (s *orderService) Create(userID, productID uint, quantity int) (*models.Order, error) {
	if quantity <= 0 {
		return nil, errors.New("quantity must be greater than zero")
	}

	// validate product exists
	_, err := s.productRepo.FindByID(productID)
	if err != nil {
		return nil, errors.New("product not found")
	}

	order := &models.Order{
		UserID:    userID,
		ProductID: productID,
		Quantity:  quantity,
	}

	if err := s.orderRepo.Create(order); err != nil {
		return nil, err
	}

	return order, nil
}

func (s *orderService) GetByUser(userID uint, limit int, offset int) ([]models.Order, int64, error) {
	return s.orderRepo.FindByUser(userID, limit, offset)
}
