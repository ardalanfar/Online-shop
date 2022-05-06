package store

import (
	"Farashop/internal/entity"
	"context"
)

//account store contract (interface)

type AccountStore interface {
	CreatedAccount(ctx context.Context, user entity.User) (entity.User, error)
	// GetAccount(ctx context.Context, userID uint) (entity.User, error)
	// UpdateAccount(ctx context.Context, user entity.User) (entity.User, error)
	// FindAccount(ctx context.Context) ([]entity.User, error)
	// DeleteAccount(ctx context.Context, userID uint) error
}
