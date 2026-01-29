package user

import "github.com/fahmiabd/go-order-api/internal/models"

type IUserRepository interface {
	Create(user *models.User) error
	FindByEmail(email string) (*models.User, error)
	FindByID(id uint) (*models.User, error)
}
