package contract

import (
	"Farashop/internal/dto"
	"context"
)

//member store(adapter) contract (interface)

type MemberStore interface {
	ShowOrders(ctx context.Context, userID uint) ([]dto.Showorders, error)
}
