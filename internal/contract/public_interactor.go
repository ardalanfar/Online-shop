package contract

import (
	"Farashop/internal/dto"
	"context"
)

//public interactor contract (interface)

type UserInteractor interface {
	Register(context.Context, dto.RegisterUserRequest) (dto.RegisterUserResponse, error)
	Login(context.Context, dto.LoginUserRequest) (dto.LoginUserResponse, error)
	MemberValidation(context.Context, dto.MemberValidationRequest) (dto.MemberValidationResponse, error)
}
