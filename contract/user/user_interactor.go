package user

import "context"

type UserInteractor interface {
	CreateUser(context.Context, dto.CreateUserRequest) (dto.CreateUserResponse, error)
	FindUser()
	FindUsers()
}
