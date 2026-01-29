package user

import (
	"errors"
	"strings"

	"github.com/fahmiabd/go-order-api/internal/model"
	"github.com/fahmiabd/go-order-api/internal/pkg/auth"
	"github.com/fahmiabd/go-order-api/internal/repositories/user"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userRepo user.IUserRepository
}

func NewUserService(userRepo user.IUserRepository) IUserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) Login(email, password string) (string, error) {
	email = strings.ToLower(strings.TrimSpace(email))
	if email == "" || password == "" {
		return "", errors.New("email and password are required")
	}

	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(password),
	)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	token, err := auth.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *userService) Register(
	name string,
	email string,
	password string,
) (*model.User, error) {
	email = strings.ToLower(strings.TrimSpace(email))

	// basic validation
	if name == "" || email == "" || password == "" {
		return nil, errors.New("name, email, and password are required")
	}

	// check existing email
	existing, err := s.userRepo.FindByEmail(email)
	if err == nil && existing != nil {
		return nil, errors.New("email already registered")
	}

	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) GetByID(id uint) (*model.User, error) {
	if id == 0 {
		return nil, errors.New("invalid user id")
	}

	return s.userRepo.FindByID(id)
}

func (s *userService) GetByEmail(email string) (*model.User, error) {
	email = strings.ToLower(strings.TrimSpace(email))
	if email == "" {
		return nil, errors.New("email is required")
	}

	return s.userRepo.FindByEmail(email)
}
