package validator

import (
	"Farashop/internal/adapter/store"
	"Farashop/internal/contract"
	"Farashop/internal/dto"
	"context"

	validation "github.com/go-ozzo/ozzo-validation"
)

func ValidateDeleteMember(store store.DbConn) contract.ValidateDeleteMember {
	return func(ctx context.Context, req dto.DeleteMemberRequest) error {
		return validation.ValidateStruct(&req,
			validation.Field(&req.ID, validation.By(DoesIDExist(ctx, store))),
		)
	}
}

func ValidateSendMessage(store store.DbConn) contract.ValidateSendMessage {
	return func(ctx context.Context, req dto.SendMsgRequest) error {
		return validation.ValidateStruct(&req,
			validation.Field(&req.ID, validation.By(DoesIDExist(ctx, store))),
		)
	}
}
