package user

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/raphaelleveque/IRGlobal/backend/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) domain.UserService {
	return &userService{repo: repo}
}

func (s *userService) Register(user *domain.User) error {
	// Verificar se o email já existe
	existingUser, err := s.repo.FindByEmail(user.Email)
	if err != nil {
		return err
	}
	if existingUser != nil {
		return errors.New("email already exists")
	}

	// Gerar hash da senha
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Preparar usuário para inserção
	user.ID = uuid.New().String()
	user.Password = string(hashedPassword)
	user.CreatedAt = time.Now()

	// Criar usuário no banco de dados
	return s.repo.Create(user)
}

// Implementar outros métodos da interface domain.UserService aqui
func (s *userService) Login(email, password string) (*domain.User, string, error) {
	return nil, "", errors.New("not implemented")
}

func (s *userService) GetAllUsers() ([]domain.User, error) {
	return nil, errors.New("not implemented")
}

func (s *userService) VerifyToken(token string) error {
	return errors.New("not implemented")
}
