package user

import "github.com/fahmiabd/go-order-api/internal/models"

type IUserService interface {
	Login(email, password string) (*models.User, error)
	Register(name, email, password string) (*models.User, error)
	GetByID(id uint) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
}
