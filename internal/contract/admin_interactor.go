package contract

import (
	"Farashop/internal/dto"
	"context"
)

//admin interactor contract (interface)

type AdminInteractor interface {
	ShowMembers(context.Context, dto.ShowMembersRequest) (dto.ShowMembersResponse, error)
	DeleteMember(context.Context, dto.DeleteMemberRequest) (dto.DeleteMemberResponse, error)
	ShowInfoMember(context.Context, dto.ShowInfoMemberRequest) (dto.ShowInfoMemberResponse, error)
}
