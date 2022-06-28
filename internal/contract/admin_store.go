package contract

import (
	"Farashop/internal/dto"
	"Farashop/internal/entity"
	"context"
)

//admin store contract (interface)

type AdminStore interface {
	ShowMembers(ctx context.Context) ([]entity.User, error)
	DeleteMember(ctx context.Context, userID uint) error
	ShowInfoMember(ctx context.Context, userID uint) (dto.ShowInfoMember, error)
}
