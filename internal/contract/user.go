package contract

import (
	"Farashop/internal/dto"
	"context"
)

//user intractor contract (interface)

type User interface {
	Register(context.Context, dto.CreateUserRequest) (dto.CreateUserResponse, error)
	//Login()
}
