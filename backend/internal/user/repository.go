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

func (r *userRepository) Create(user *domain.User) error {
	query := `
		INSERT INTO users (id, email, password, name, created_at)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := r.db.Exec(query, user.ID, user.Email, user.Password, user.Name, user.CreatedAt)
	return err
}

func (r *userRepository) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	query := `SELECT id, email, password, name, created_at FROM users WHERE email = $1`

	err := r.db.QueryRow(query, email).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.Name,
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
