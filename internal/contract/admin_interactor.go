package contract

import (
	"Farashop/internal/dto/admin"
	"context"
)

//admin contract (interface)

type AdminInteractor interface {
	ShowMembers(context.Context, admin.ShowMembersRequest) (admin.ShowMembersResponse, error)
	DeleteMember(context.Context, admin.DeleteMemberRequest) (admin.DeleteMemberRequest, error)
}
