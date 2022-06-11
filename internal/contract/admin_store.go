package contract

import (
	"Farashop/internal/entity"
	"context"
)

//admin store(adapter) contract (interface)

type AdminStore interface {
	ShowMembers(ctx context.Context) ([]entity.User, error)
	DeleteMember(ctx context.Context, user entity.User) error
	ShowInfoMember(ctx context.Context, user entity.User) (entity.User, error)
}
