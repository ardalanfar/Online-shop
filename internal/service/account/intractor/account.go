package intractor

//Interactor Account (loign , register)

import (
	psql "Farashop/internal/repository/postgresql"
)

type Interactor struct {
	store psql.AccountStore
}

func New(store psql.AccountStore) Interactor {
	return Interactor{store: store}
}

// func (i Interactor) login(ctx context.Context, req account.LoginAccountRequest) (account.LoginAccountResponse, error) {

// }

// func (i Interactor) Register(ctx context.Context, req account.RegisterAccountRequest) error {

// }
