package contract

import (
	"Farashop/internal/dto/public_dto"
	"context"
)

//user contract (interface)

type UserInteractor interface {
	Register(context.Context, public_dto.CreateUserRequest) (public_dto.CreateUserResponse, error)
	Login(context.Context, public_dto.LoginUserRequest) (public_dto.CreateUserResponse, error)
}
