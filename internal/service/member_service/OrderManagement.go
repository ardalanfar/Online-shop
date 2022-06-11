package member_service

import (
	"Farashop/internal/contract"
	"Farashop/internal/dto"
	"context"
)

type Interactor struct {
	store contract.MemberStore
}

func NewMember(store contract.MemberStore) Interactor {
	return Interactor{store: store}
}

func (i Interactor) ShowOrders(ctx context.Context, _ dto.ShowOrdersRequest) (dto.ShowOrdersResponse, error) {

	//Show Orders
	orders, err := i.store.ShowOrders(ctx)
	if err != nil {
		return dto.ShowOrdersResponse{}, err
	}
	//return
	return dto.ShowOrdersResponse{Orders: orders}, nil
}
