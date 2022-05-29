package admin_service

//Admin intractor (ShowMembers , )

import (
	"Farashop/internal/contract"
	"Farashop/internal/dto/admin_dto"
	"context"
)

type Interactor struct {
	store contract.AdminStore
}

func New(store contract.AdminStore) Interactor {
	return Interactor{store: store}
}

func (i Interactor) ShowMembers(ctx context.Context, _ admin_dto.ShowMembersRequest) (admin_dto.ShowMembersResponse, error) {
	users, err := i.store.ShowMembers(ctx)
	if err != nil {
		return admin_dto.ShowMembersResponse{}, err
	}

	return admin_dto.ShowMembersResponse{Users: users}, nil
}

// func (i Interactor) DeleteMember(ctx context.Context, req admin.DeleteMemberRequest) (admin.DeleteMemberResponse, error) {

// }
