package domain

import (
	"time"
)

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
}

type UserRepository interface {
	Create(user *User) (*User, error)
	FindByEmail(email string) (*User, error)
	FindByID(id string) (*User, error)

}

type UserService interface {
	Register(user *User) (*User, error)
	GetByID(id string) (*User, error)
	GetByEmail(email string) (*User, error)
}