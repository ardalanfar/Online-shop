package store

//Account Store intractor

import (
	"Farashop/internal/adapter/store/model"
	"Farashop/internal/entity"
	"context"
)

func (s DbConn) CreateUser(ctx context.Context, user entity.User) (entity.User, error) {

	u := model.MapFromUserEntity(user)
	err := s.Db.WithContext(ctx).Create(&u).Error

	if err != nil {
		return entity.User{}, err
	}

	return model.MapToUserEntity(u), nil
}
func (s DbConn) GetUserByUsername(ctx context.Context, user entity.User) (entity.User, error) {

	u := model.MapFromUserEntity(user)
	err := s.Db.WithContext(ctx).Select("id", "password", "username").First(&u)

	if err != nil {
		return entity.User{}, err.Error
	}

	return model.MapToUserEntity(u), nil
}
