package user

import (
	"database/sql"

	"github.com/raphaelleveque/IRGlobal/backend/internal/domain"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) domain.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *domain.User) (*domain.User, error) {
	query := `
		INSERT INTO users (name, email, password)
		VALUES ($1, $2, $3)
		RETURNING id, name, email, password, created_at
	`
	var response domain.User
	err := r.db.QueryRow(query, user.Name, user.Email, user.Password).Scan(
		&response.ID,
		&response.Name,
		&response.Email,
		&response.Password,
		&response.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (r *userRepository) FindByEmail(email string) (*domain.User, error) {
	query := `
		SELECT id, name, email, password, created_at FROM users
		WHERE email = $1
	`
	var user domain.User
	err := r.db.QueryRow(query, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByID(id string) (*domain.User, error) {
	query := `
		SELECT id, name, email, password, created_at FROM users
		WHERE id = $1
	`
	var user domain.User
	err := r.db.QueryRow(query, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}
