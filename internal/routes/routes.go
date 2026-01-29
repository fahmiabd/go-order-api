package routes

import (
	"github.com/go-chi/chi/v5"

	authController "github.com/fahmiabd/go-order-api/internal/controllers/auth"
	orderController "github.com/fahmiabd/go-order-api/internal/controllers/order"
	"github.com/fahmiabd/go-order-api/internal/middleware"
	"github.com/fahmiabd/go-order-api/internal/services/auth"
	"github.com/fahmiabd/go-order-api/internal/services/order"
	"github.com/fahmiabd/go-order-api/internal/services/user"
)

type RouterDeps struct {
	AuthService  auth.IAuthService
	UserService  user.IUserService
	OrderService order.IOrderService
}

func Register(r chi.Router, deps RouterDeps) {
	authController := authController.NewAuthController(
		deps.UserService,
		deps.AuthService,
	)

	orderController := orderController.NewOrderController(
		deps.OrderService,
	)

	// public routes
	r.Post("/auth/login", authController.Login)
	r.Post("/auth/register", authController.Register)

	// protected routes
	r.Route("/orders", func(r chi.Router) {
		r.Use(middleware.AuthMiddleware(deps.AuthService))
		r.Post("/", orderController.Create)
	})
}
