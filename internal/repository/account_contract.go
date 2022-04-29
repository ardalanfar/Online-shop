package repository

import (
	"Farashop/internal/entity"
	"context"
)

//Account Store contract (interface)

type AccountStore interface {
	createdAccount(ctx context.Context, user entity.User) (entity.User, error)
	GetAccount(ctx context.Context, userID uint) (entity.User, error)
	UpdateAccount(ctx context.Context, user entity.User) (entity.User, error)
	FindAccount(ctx context.Context) ([]entity.User, error)
	DeleteAccount(ctx context.Context, userID uint) error
}
