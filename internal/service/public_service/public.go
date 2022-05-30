package public_service

//User intractor (loign , register)

import (
	"Farashop/internal/contract"
	"Farashop/internal/dto"
	"Farashop/internal/entity"
	"Farashop/pkg/encrypt"
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

	//create hash password
	Password, errhash := encrypt.HashPassword(user.Password)
	if errhash != nil {
		return dto.CreateUserResponse{Result: false}, errhash
	}
	user.Password = Password

	//create user
	createdUser, errCrate := i.store.Register(ctx, user)
	if errCrate != nil {
		return dto.CreateUserResponse{Result: false}, errCrate
	}
	//return
	return dto.CreateUserResponse{User: createdUser}, nil
}

func (i Interactor) Login(ctx context.Context, req dto.LoginUserRequest) (dto.LoginUserResponse, error) {
	user := entity.User{
		Username: req.Username,
		Password: req.Password,
	}

	//get information user by username
	getInfo, errInfo := i.store.GetUserByUsername(ctx, user)
	if errInfo != nil {
		return dto.LoginUserResponse{}, errInfo
	}
	//check password with username
	checkpass := encrypt.CheckPasswordHash(user.Password, getInfo.Password)
	if checkpass != nil {
		return dto.LoginUserResponse{Result: false, User: getInfo}, checkpass
	}
	//return
	return dto.LoginUserResponse{Result: true, User: getInfo}, nil
}
