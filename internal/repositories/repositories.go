package repositories

import (
	"github.com/fahmiabd/go-order-api/internal/repositories/order"
	"github.com/fahmiabd/go-order-api/internal/repositories/product"
	"github.com/fahmiabd/go-order-api/internal/repositories/user"
	"gorm.io/gorm"
)

type Repositories struct {
	OrderRepo   order.IOrderRepository
	ProductRepo product.IProductRepository
	UserRepo    user.IUserRepository
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		OrderRepo:   order.NewOrderRepository(db),
		ProductRepo: product.NewProductRepository(db),
		UserRepo:    user.NewUserRepository(db),
	}
}
