package app

import (
	"context"
	"errors"
)

var ErrUserNotFound = errors.New("user is not found")

type User struct {
	ID   int
	Name string
}

type UserRepository interface {
	GetByID(ctx context.Context, id int) (User, error)
}

type UserController interface {
	GetUser(ctx context.Context, id int) (User, error)
}
