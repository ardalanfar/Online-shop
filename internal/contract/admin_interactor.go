package contract

import (
	"Farashop/internal/dto/admin_dto"
	"context"
)

//admin contract (interface)

type AdminInteractor interface {
	ShowMembers(context.Context, admin_dto.ShowMembersRequest) (admin_dto.ShowMembersResponse, error)
	//DeleteMember(context.Context, admin_dto.DeleteMemberRequest) (admin_dto.DeleteMemberRequest, error)
}
