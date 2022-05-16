package contract

import (
	"Farashop/internal/dto"
)

type (
	ValidateCreateUser func(req dto.CreateUserRequest) error
	ValidateLoginUser  func(req dto.LoginUserRequest) error
)