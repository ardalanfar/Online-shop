package postgresql

import (
	"Farashop/internal/entity"
	"context"
)

//Account Store contract (interface)

type AccountStore interface {
	CreateUser(ctx context.Context, user entity.User) (entity.User, error)
	GetUser(ctx context.Context, userID uint) (entity.User, error)
	UpdateUser(ctx context.Context, user entity.User) (entity.User, error)
	FindUsers(ctx context.Context) ([]entity.User, error)
	DeleteUser(ctx context.Context, userID uint) error
}
