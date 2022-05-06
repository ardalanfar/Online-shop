package account

//Account intractor (loign , register)

import (
	"Farashop/internal/entity"
	"Farashop/internal/repository"
	"context"
)

type Interactor struct {
	store repository.AccountStore
}

func New(store repository.AccountStore) Interactor {
	return Interactor{store: store}
}

func (i Interactor) Register(ctx context.Context, req RegisterAccountRequest) (RegisterAccountResponse, error) {

	user := entity.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	//Password := pkg.Hash.HashPassword(user.Password)
	//user.Password = Password

	createdAccount, err := i.store.CreatedAccount(ctx, user)
	if err != nil {
		return RegisterAccountResponse{}, err
	}

	return RegisterAccountResponse{User: createdAccount}, nil
}

// func (i Interactor) login(ctx context.Context, req LoginAccountRequest) (LoginAccountResponse, error) {

// }
