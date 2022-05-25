package model

import "Farashop/internal/entity"

//user model database
type User struct {
	ID        uint   `json:"id" gorm:"primary_key,serializer:json"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	ID_access uint   `json:"id_access"`
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
