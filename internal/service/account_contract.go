package service

import (
	"Farashop/internal/service/account"
	"context"
)

//Account contract (interface)

type Account interface {
	Login(context.Context, account.LoginAccountRequest) (account.LoginAccountResponse, error)
	Register(context.Context, account.RegisterAccountRequest) (LoginAccountResponse error)
}
