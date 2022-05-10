package user

//User intractor (loign , register)

import (
	"Farashop/internal/adapter/store"
	"Farashop/internal/dto"
	"Farashop/internal/entity"
	"Farashop/internal/pkg"
	"context"
)

type Interactor struct {
	store store.DbConn
}

func New(store store.DbConn) Interactor {
	return Interactor{store: store}
}

func (i Interactor) Register(ctx context.Context, req dto.CreateUserRequest) (dto.CreateUserResponse, error) {
	user := entity.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	//Password := pkg.Hash.HashPassword(user.Password)
	//user.Password = Password

	createdUser, err := i.store.CreateUser(ctx, user)
	if err != nil {
		return dto.CreateUserResponse{}, err
	}

	return dto.CreateUserResponse{User: createdUser}, nil
}

func (i Interactor) Login(ctx context.Context, req dto.LoginUserRequest) (dto.LoginUserResponse, error) {
	user := entity.User{
		Username: req.Username,
	}

	getUsername, err := i.store.GetUserByUsername(ctx, user)

	if err != nil {
		return dto.LoginUserResponse{}, err
	}

	pold := getUsername.Password
	pnew := pkg.Hash.DecodePassword(pold)

}
