package domain

import "errors"

var (
	ErrInsufficientQuantity = errors.New("you cannot sell more quantities than what you have")
)
