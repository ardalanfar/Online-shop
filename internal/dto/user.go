package dto

import "Farashop/internal/entity"

//User Contract Interactor Request And Response (Struct)

/*-----------------Create----------------------*/

type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserResponse struct {
	User entity.User `json:"user"`
}

/*--------------------Login----------------------*/

type LoginUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginUserResponse struct {
	Result bool        `json:"result"`
	User   entity.User `json:"user"`
}

/*-----------------------------------------------*/
