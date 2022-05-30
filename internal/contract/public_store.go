package contract

import (
	"Farashop/internal/entity"
	"context"
)

//Public store contract (interface)

type UserStore interface {
	Register(ctx context.Context, user entity.User) (entity.User, error)
	GetUserByUsername(ctx context.Context, user entity.User) (entity.User, error)
}
