package adminservice

import (
	"Farashop/internal/contract"
	"Farashop/internal/dto/admin"
	"context"
)

//Admin intractor (ShowMembers , )

type Interactor struct {
	store contract.AdminStore
}

func New(store contract.AdminStore) Interactor {
	return Interactor{store: store}
}

func (i Interactor) ShowMembers(ctx context.Context, req admin.ShowMembersRequest) (admin.ShowMembersResponse, error) {
	users, err := i.store.ShowMembers(ctx)
	if err != nil {
		return admin.ShowMembersResponse{}, err
	}

	return admin.ShowMembersResponse{Users: users}, nil
}

// func (i Interactor) DeleteMember(ctx context.Context, req admin.DeleteMemberRequest) (admin.DeleteMemberResponse, error) {

// }
