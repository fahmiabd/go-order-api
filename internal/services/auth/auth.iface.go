package auth

type IAuthService interface {
	GenerateToken(userID uint) (string, error)
	ValidateToken(token string) (uint, error)
}
