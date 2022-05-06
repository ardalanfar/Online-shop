package model

import "Farashop/internal/entity"

//User model database

type User struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func MapFromUserEntity(user entity.User) User {
	return User{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}
}

func MapToUserEntity(user User) entity.User {
	return entity.User{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}
}
