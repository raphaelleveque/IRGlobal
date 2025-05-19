package domain

import "time"

type User struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type UserRepository interface {
	Create(user *User) error
	FindByEmail(email string) (*User, error)
}

type UserService interface {
	Register(user *User) error
	Login(email, password string) (*User, string, error)
	GetAllUsers() ([]User, error)
	VerifyToken(token string) error
}
