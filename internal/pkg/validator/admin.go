package validator

import (
	"Farashop/internal/adapter/store"
	"Farashop/internal/contract"
	"Farashop/internal/dto/admin_dto"
	"context"

	validation "github.com/go-ozzo/ozzo-validation"
)

func ValidateDeleteMember(store store.DbConn) contract.ValidateDeleteMember {
	return func(ctx context.Context, req admin_dto.DeleteMemberRequest) error {
		return validation.ValidateStruct(&req,
			validation.Field(&req.ID, validation.Required),
		)
	}
}
