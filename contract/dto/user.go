package user

import "restful/entity"

type CreateUserRequest struct {
	Fullname string `json:"fullname"`
	Age      int    `json:"age"`
	Courseid int    `json:"courseid"`
}

type CreateUserResponse struct {
	user entity.User `json:"user"`
}

type FindUsersRequest struct{}

type FindUsersResponse struct {
	Users []entity.User `json:"users"`
}

type FindUserRequest struct {
	Id uint `json:"id"`
}

type FindUserResponse struct {
	User entity.User `json:"user"`
}
