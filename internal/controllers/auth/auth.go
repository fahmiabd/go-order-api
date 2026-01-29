package auth

import (
	"encoding/json"
	"net/http"

	"github.com/fahmiabd/go-order-api/internal/dto"
	authService "github.com/fahmiabd/go-order-api/internal/services/auth"
	userService "github.com/fahmiabd/go-order-api/internal/services/user"
)

type AuthController struct {
	userService userService.IUserService
	authService authService.IAuthService
}

func NewAuthController(userService userService.IUserService, authService authService.IAuthService) *AuthController {
	return &AuthController{
		userService: userService,
		authService: authService,
	}
}

func (h *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	var req dto.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	user, err := h.userService.Login(req.Email, req.Password)
	if err != nil {
		http.Error(w, "invalid email or password", http.StatusUnauthorized)
		return
	}

	token, err := h.authService.GenerateToken(user.ID)
	if err != nil {
		http.Error(w, "failed to generate token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}

func (h *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	var req dto.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	user, err := h.userService.Register(req.Name, req.Email, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
