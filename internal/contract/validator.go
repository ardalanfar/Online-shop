package contract

import (
	"Farashop/internal/dto"
)

//validator Contract (interface)

type (
	ValidateCreateUser func(req dto.CreateUserRequest) error
	ValidateLoginUser  func(req dto.LoginUserRequest) error
)
