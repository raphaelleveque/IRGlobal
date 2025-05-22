package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/raphaelleveque/IRGlobal/backend/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	secretKey   []byte
	userService domain.UserService
}

func NewAuthService(secretKey []byte, userService domain.UserService) domain.AuthService {
	return &authService{secretKey: secretKey, userService: userService}
}

func (s *authService) GenerateToken(user *domain.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(s.secretKey)
	if err != nil {
		return "", nil
	}

	return tokenString, nil
}

func (s *authService) ValidateToken(tokenString string) (*domain.User, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("método de assinatura inválido")
		}
		return s.secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userIDStr, ok := claims["user_id"].(string)
		if _, err := uuid.Parse(userIDStr); !ok || err != nil {
			return nil, errors.New("user_id inválido no token")
		}
		return s.userService.GetByID(userIDStr)
	}

	return nil, errors.New("token inválido")
}

// ValidatePassword compara a senha fornecida com o hash armazenado usando bcrypt
func (s *authService) ValidatePassword(plainPassword, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}

func (s *authService) GeneratePasswordHash(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
