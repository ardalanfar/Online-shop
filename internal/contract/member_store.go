package contract

import (
	"Farashop/internal/entity"
	"context"
)

//member store(adapter) contract (interface)

type MemberStore interface {
	ShowOrders(ctx context.Context, userID uint) ([]entity.Order, error)
}
