package contract

import (
	"Farashop/internal/dto"
	"context"
)

//member interactor contract (interface)

type MemberInteractor interface {
	ShowOrders(context.Context, dto.ShowOrdersRequest) (dto.ShowOrdersResponse, error)
}
