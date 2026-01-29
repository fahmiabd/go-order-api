package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"

	"github.com/fahmiabd/go-order-api/internal/config"
	"github.com/fahmiabd/go-order-api/internal/models"
	"github.com/fahmiabd/go-order-api/internal/repositories"
	"github.com/fahmiabd/go-order-api/internal/routes"
	"github.com/fahmiabd/go-order-api/internal/services"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	r := chi.NewRouter()

	// basic middlewares
	r.Use(middleware.Recoverer)

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

	repositories := repositories.NewRepositories(db)
	services := services.NewServices(repositories)

	// ===== routes =====
	routes.Register(r, routes.RouterDeps{
		AuthService:  services.AuthService,
		UserService:  services.UserService,
		OrderService: services.OrderService,
	})

	log.Println("server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
