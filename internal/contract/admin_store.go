package contract

import (
	"Farashop/internal/entity"
	"context"
)

//admin store contract (interface)

type AdminStore interface {
	ShowMembers(ctx context.Context) ([]entity.User, error)
	//DeleteMember(ctx context.Context, userId uint) error
}
