package contract

import (
	"Farashop/internal/entity"
	"context"
)

//user store contract (interface)

type UserStore interface {
	CreateUser(ctx context.Context, user entity.User) (entity.User, error)
	//EditUser()
}
