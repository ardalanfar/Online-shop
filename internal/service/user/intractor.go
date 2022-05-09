package user

//User intractor (loign , register)

import (
	"Farashop/internal/contract"
	"Farashop/internal/dto"
	"Farashop/internal/entity"
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

	//Password := pkg.Hash.HashPassword(user.Password)
	//user.Password = Password

	createdAccount, err := i.store.CreateUser(ctx, user)
	if err != nil {
		return dto.CreateUserResponse{}, err
	}

	return dto.CreateUserResponse{User: createdAccount}, nil
}

// func (i Interactor) login(ctx context.Context, req LoginAccountRequest) (LoginAccountResponse, error) {
// }
