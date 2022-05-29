package public_dto

import "Farashop/internal/entity"

//User Contract Interactor Request And Response (Struct)

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
