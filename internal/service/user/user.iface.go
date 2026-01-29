package user

import "github.com/fahmiabd/go-order-api/internal/model"

type IUserService interface {
	Login(email, password string) (string, error)
	Register(name, email, password string) (*model.User, error)
	GetByID(id uint) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
}
