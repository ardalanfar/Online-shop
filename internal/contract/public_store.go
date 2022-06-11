package contract

import (
	"Farashop/internal/entity"
	"context"
)

//public store(adapter) contract (interface)

type PublicStore interface {
	Register(ctx context.Context, user entity.User) error
	Login(ctx context.Context, user entity.User) (entity.User, error)
	MemberValidation(ctx context.Context, user entity.User) (bool, error)
}
