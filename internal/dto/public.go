package dto

//Public Contract Interactor Request And Response (Struct)

import "Farashop/internal/entity"

/*-----------------Create user----------------------*/

type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserResponse struct {
	Result bool        `json:"result"`
	User   entity.User `json:"user"`
}

/*--------------------Login user----------------------*/

type LoginUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginUserResponse struct {
	Result bool        `json:"result"`
	User   entity.User `json:"user"`
}

/*-----------------------------------------------------*/
