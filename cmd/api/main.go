package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"

	"github.com/fahmiabd/go-order-api/internal/config"
	"github.com/fahmiabd/go-order-api/internal/models"
	"github.com/fahmiabd/go-order-api/internal/pkg/auth"
	orderRepositories "github.com/fahmiabd/go-order-api/internal/repositories/order"
	productRepositories "github.com/fahmiabd/go-order-api/internal/repositories/product"
	userRepositories "github.com/fahmiabd/go-order-api/internal/repositories/user"
	"github.com/fahmiabd/go-order-api/internal/routes"
	authService "github.com/fahmiabd/go-order-api/internal/services/auth"
	orderService "github.com/fahmiabd/go-order-api/internal/services/order"
	userService "github.com/fahmiabd/go-order-api/internal/services/user"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	r := chi.NewRouter()

	// basic middlewares
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// ===== dependencies =====
	jwtManager := auth.NewJWTManager(
		os.Getenv("JWT_SECRET"),
		24*time.Hour,
	)

	db, err := config.InitDB()
	if err != nil {
		log.Println("warning: database not connected, running without DB")
	}

	// ===== auto migrate =====
	if err := db.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.Order{},
	); err != nil {
		log.Fatal("failed to auto migrate:", err)
	}

	userRepo := userRepositories.NewUserRepository(db)
	orderRepo := orderRepositories.NewOrderRepository(db)
	productRepo := productRepositories.NewProductRepository(db)

	authService := authService.NewAuthService(jwtManager)
	userService := userService.NewUserService(userRepo)
	orderService := orderService.NewOrderService(orderRepo, productRepo)

	// ===== routes =====
	routes.Register(r, routes.RouterDeps{
		AuthService:  authService,
		UserService:  userService,
		OrderService: orderService,
	})

	log.Println("server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
