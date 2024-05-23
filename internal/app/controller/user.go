package controller

import (
	"context"
	"errors"

	"github.com/a-reshatniak-itechart/projectlayout/internal/app"
)

func NewUser(repo app.UserRepository) User {
	return User{repo: repo}
}

type User struct {
	repo app.UserRepository
}

func (c User) GetUser(ctx context.Context, id int) (app.User, error) {
	if id <= 0 {
		return app.User{}, errors.New("invalid user ID")
	}

	return c.repo.GetByID(ctx, id)
}
