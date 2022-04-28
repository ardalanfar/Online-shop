package account

import "Farashop/internal/entity"

//Account Contract Request And Response (Struct)

type LoginAccountRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginAccountResponse struct {
	User entity.User `json:"user"`
}

type RegisterAccountRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterAccountResponse struct {
	User entity.User `json:"user"`
}
