package service

import (
	"Farashop/internal/dto"
)

type (
	ValidateCreateUser func(req dto.CreateUserRequest) error
)
