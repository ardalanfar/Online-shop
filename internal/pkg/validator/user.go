package validator

import (
	"Farashop/internal/adapter/store"
	"Farashop/internal/contract"
	"Farashop/internal/dto"
	"context"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

//Intractor package validator

func ValidateCreateUser(store store.DbConn) contract.ValidateCreateUser {
	return func(ctx context.Context, req dto.CreateUserRequest) error {
		return validation.ValidateStruct(&req,
			validation.Field(&req.Username, validation.By(DoesUsernameExist(ctx, store))),
			validation.Field(&req.Email, validation.Required, is.Email),
			validation.Field(&req.Password, validation.Required),
		)
	}
}

func ValidateLoginUser(store store.DbConn) contract.ValidateLoginUser {
	return func(ctx context.Context, req dto.LoginUserRequest) error {
		return validation.ValidateStruct(&req,
			validation.Field(&req.Username, validation.By(DoesUsernameExist(ctx, store))),
			validation.Field(&req.Password, validation.Required),
		)
	}
}
