package contract

import (
	"Farashop/internal/dto"
	"context"
)

//validator interactor package Contract

type (
	ValidateRegister         func(ctx context.Context, req dto.RegisterUserRequest) error
	ValidateLogin            func(ctx context.Context, req dto.LoginUserRequest) error
	ValidateDeleteMember     func(ctx context.Context, req dto.DeleteMemberRequest) error
	ValidateShowInfoMember   func(ctx context.Context, req dto.ShowInfoMemberRequest) error
	ValidateMemberValidation func(ctx context.Context, req dto.MemberValidationRequest) error
)
