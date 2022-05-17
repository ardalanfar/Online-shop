package user

//User intractor (loign , register)

import (
	"Farashop/internal/contract"
	"Farashop/internal/dto"
	"Farashop/internal/entity"
	"Farashop/internal/pkg/hash"
	"context"
)

type Interactor struct {
	store contract.UserStore
}

func New(store contract.UserStore) Interactor {
	return Interactor{store: store}
}

func (i Interactor) Register(ctx context.Context, req dto.CreateUserRequest) (dto.CreateUserResponse, error) {
	user := entity.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	Password, errhash := hash.HashPassword(user.Password)
	if errhash != nil {
		return dto.CreateUserResponse{}, errhash
	}
	user.Password = Password

	createdUser, err := i.store.CreateUser(ctx, user)
	if err != nil {
		return dto.CreateUserResponse{}, err
	}

	return dto.CreateUserResponse{User: createdUser}, nil
}

func (i Interactor) Login(ctx context.Context, req dto.LoginUserRequest) (dto.LoginUserResponse, error) {
	user := entity.User{
		Username: req.Username,
		Password: req.Password,
	}

	getInfo, errInfo := i.store.GetUserByUsername(ctx, user)

	if errInfo != nil {
		return dto.LoginUserResponse{}, errInfo
	}

	res := hash.CheckPasswordHash(user.Password, getInfo.Password)

	return dto.LoginUserResponse{Result: res}, nil
}
