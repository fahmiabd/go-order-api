package services

import (
	"os"
	"time"

	pkgAuth "github.com/fahmiabd/go-order-api/internal/pkg/auth"
	"github.com/fahmiabd/go-order-api/internal/repositories"
	"github.com/fahmiabd/go-order-api/internal/services/auth"
	"github.com/fahmiabd/go-order-api/internal/services/order"
	"github.com/fahmiabd/go-order-api/internal/services/user"
)

type Services struct {
	AuthService  auth.IAuthService
	OrderService order.IOrderService
	UserService  user.IUserService
}

func NewServices(repositories *repositories.Repositories) *Services {
	jwtManager := pkgAuth.NewJWTManager(
		os.Getenv("JWT_SECRET"),
		24*time.Hour,
	)

	return &Services{
		AuthService:  auth.NewAuthService(jwtManager),
		OrderService: order.NewOrderService(repositories.OrderRepo, repositories.ProductRepo),
		UserService:  user.NewUserService(repositories.UserRepo),
	}
}
