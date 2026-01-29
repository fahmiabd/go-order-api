package auth

import (
	"errors"

	"github.com/fahmiabd/go-order-api/internal/pkg/auth"
)

type authService struct {
	jwt *auth.JWTManager
}

func NewAuthService(jwt *auth.JWTManager) IAuthService {
	return &authService{
		jwt: jwt,
	}
}

func (s *authService) GenerateToken(userID uint) (string, error) {
	return s.jwt.Generate(userID)
}

func (s *authService) ValidateToken(token string) (uint, error) {
	claims, err := s.jwt.Parse(token)
	if err != nil {
		return 0, errors.New("invalid token")
	}
	return claims.UserID, nil
}
