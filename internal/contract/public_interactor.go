package contract

import (
	"Farashop/internal/dto"
	"context"
)

//Public interactor contract (interface)

type UserInteractor interface {
	Register(context.Context, dto.CreateUserRequest) (dto.CreateUserResponse, error)
	Login(context.Context, dto.LoginUserRequest) (dto.LoginUserResponse, error)
}
