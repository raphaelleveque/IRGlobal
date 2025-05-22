package user

import (
	"errors"

	"github.com/raphaelleveque/IRGlobal/backend/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) domain.UserService {
	return &userService{repo: repo}
}

func (s *userService) Register(user *domain.User) (*domain.User, error) {
	existingUser, err := s.repo.FindByEmail(user.Email)
	if err != nil {
		return nil, err
	}

	if existingUser != nil {
		return nil, errors.New("email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashedPassword)

	return s.repo.Create(user)
}

func (s *userService) GetByID(id string) (*domain.User, error) {
	existingUser, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return existingUser, nil
}

func (s *userService) GetByEmail(email string) (*domain.User, error) {
	existingUser, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	return existingUser, nil
}