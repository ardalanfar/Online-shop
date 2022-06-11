package contract

import (
	"Farashop/internal/entity"
	"context"
)

//member store(adapter) contract (interface)

type MemberStore interface {
	ShowOrders(ctx context.Context) ([]entity.Order, error)
}
