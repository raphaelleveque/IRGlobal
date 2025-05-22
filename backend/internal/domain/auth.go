package domain

type AuthService interface {
	ValidateToken(tokenString string) (*User, error)
	GenerateToken(user *User) (string, error)
}
