package user

import "github.com/fahmiabd/go-order-api/internal/model"

type IUserRepository interface {
	Create(user *model.User) error
	FindByEmail(email string) (*model.User, error)
	FindByID(id uint) (*model.User, error)
}
