package contract

import (
	"Farashop/internal/dto"
	"context"
)

//Validator interactor package Contract (interface)

type (
	ValidateCreateUser   func(ctx context.Context, req dto.CreateUserRequest) error
	ValidateLoginUser    func(ctx context.Context, req dto.LoginUserRequest) error
	ValidateDeleteMember func(ctx context.Context, req dto.DeleteMemberRequest) error
	ValidateSendMessage  func(ctx context.Context, req dto.SendMsgRequest) error
)
