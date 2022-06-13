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

	id := uint(2)
	//show orders
	getInfo, errInfo := i.store.ShowOrders(ctx, id)
	if errInfo != nil {
		return dto.ShowOrdersResponse{}, errInfo
	}
	//return
	return dto.ShowOrdersResponse{Orders: getInfo}, nil
}
