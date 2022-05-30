package admin_service

//Admin intractor (ShowMembers , )

import (
	"Farashop/internal/contract"
	"Farashop/internal/dto"
	"Farashop/internal/entity"
	"context"
)

type Interactor struct {
	store contract.AdminStore
}

func New(store contract.AdminStore) Interactor {
	return Interactor{store: store}
}

func (i Interactor) ShowMembers(ctx context.Context, _ dto.ShowMembersRequest) (dto.ShowMembersResponse, error) {
	users, err := i.store.ShowMembers(ctx)
	if err != nil {
		return dto.ShowMembersResponse{}, err
	}
	//return
	return dto.ShowMembersResponse{Users: users}, nil
}

func (i Interactor) DeleteMember(ctx context.Context, req dto.DeleteMemberRequest) (dto.DeleteMemberResponse, error) {
	user := entity.User{
		ID: req.ID,
	}

	//get information user by username
	errInfo := i.store.DeleteMember(ctx, user)
	if errInfo != nil {
		return dto.DeleteMemberResponse{}, errInfo
	}
	//return
	return dto.DeleteMemberResponse{Result: true}, nil
}
