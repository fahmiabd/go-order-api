package order

import (
	"encoding/json"
	"net/http"

	"github.com/fahmiabd/go-order-api/internal/dto"
	"github.com/fahmiabd/go-order-api/internal/middleware"
	orderService "github.com/fahmiabd/go-order-api/internal/services/order"
)

type OrderController struct {
	orderService orderService.IOrderService
}

func NewOrderController(orderService orderService.IOrderService) *OrderController {
	return &OrderController{
		orderService: orderService,
	}
}

func (c *OrderController) Create(w http.ResponseWriter, r *http.Request) {
	// ambil user_id dari context (hasil auth middleware)
	userID := middleware.UserIDFromContext(r.Context())
	if userID < 1 {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	var req dto.CreateOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if req.ProductID == 0 || req.Quantity <= 0 {
		http.Error(w, "product_id and quantity are required", http.StatusBadRequest)
		return
	}

	order, err := c.orderService.Create(
		userID,
		req.ProductID,
		req.Quantity,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}
