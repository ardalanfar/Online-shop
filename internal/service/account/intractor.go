package account

//Account intractor (loign , register)

import (
	"Farashop/internal/entity"
	"Farashop/internal/pkg"
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

	user.Password = pkg.Hash.HashPassword(user.Password)

	createdAccount, err := i.store.createdAccount(ctx, user)

	//return
}

// func (i Interactor) login(ctx context.Context, req LoginAccountRequest) (LoginAccountResponse, error) {

// }
