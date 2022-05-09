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

/*-----------------------------------------------*/
