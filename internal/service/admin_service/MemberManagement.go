package admin_service

import (
	"Farashop/internal/contract"
	"Farashop/internal/dto"
	"context"
)

type Interactor struct {
	store contract.AdminStore
}

func NewAdmin(store contract.AdminStore) Interactor {
	return Interactor{store: store}
}

func (i Interactor) ShowMembers(ctx context.Context, _ dto.ShowMembersRequest) (dto.ShowMembersResponse, error) {
	//Show Members
	users, err := i.store.ShowMembers(ctx)
	if err != nil {
		return dto.ShowMembersResponse{}, err
	}
	//return
	return dto.ShowMembersResponse{Users: users}, nil
}

func (i Interactor) DeleteMember(ctx context.Context, req dto.DeleteMemberRequest) (dto.DeleteMemberResponse, error) {

	//Delete Member
	errInfo := i.store.DeleteMember(ctx, req.ID)
	if errInfo != nil {
		return dto.DeleteMemberResponse{Result: false}, errInfo
	}
	//return true
	return dto.DeleteMemberResponse{Result: true}, nil
}

func (i Interactor) ShowInfoMember(ctx context.Context, req dto.ShowInfoMemberRequest) (dto.ShowInfoMemberResponse, error) {

	//get information user by username
	getInfo, errInfo := i.store.ShowInfoMember(ctx, req.ID)
	if errInfo != nil {
		return dto.ShowInfoMemberResponse{User: getInfo}, errInfo
	}
	//return true
	return dto.ShowInfoMemberResponse{User: getInfo}, nil

}
