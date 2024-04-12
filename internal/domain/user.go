package domain

import (
	"context"
	"time"
)

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"title" validate:"required"`
	Login     string    `json:"login" validate:"required"`
	Password  string    `json:"password" validate:"required"`
	UpdatedAt time.Time `json:"updatedAt"`
	CreatedAt time.Time `json:"createdAt"`
}

type UserUseCase interface {
	Create(ctx context.Context, cursor interface{}) ([]User, string, error)
}

type UserRepository interface {
	Create(ctx context.Context, cursor interface{}) (res []User, err error)
}
