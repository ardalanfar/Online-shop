package contract

import (
	"Farashop/internal/dto"
	"context"
)

//validator Contract (interface)

type (
	ValidateCreateUser func(ctx context.Context, req dto.CreateUserRequest) error
	ValidateLoginUser  func(ctx context.Context, req dto.LoginUserRequest) error
)
