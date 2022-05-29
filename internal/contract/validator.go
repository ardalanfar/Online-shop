package contract

import (
	"Farashop/internal/dto/admin_dto"
	"Farashop/internal/dto/public_dto"
	"context"
)

//validator Contract (interface)

type (
	ValidateCreateUser   func(ctx context.Context, req public_dto.CreateUserRequest) error
	ValidateLoginUser    func(ctx context.Context, req public_dto.LoginUserRequest) error
	ValidateDeleteMember func(ctx context.Context, req admin_dto.DeleteMemberRequest) error
)
