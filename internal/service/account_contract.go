package service

import (
	"Farashop/internal/dto"
	"context"
)

//Account contract (interface)

type Account interface {
	Register(context.Context, dto.CreateUserRequest) (dto.CreateUserResponse, error)
	//Login()
}
