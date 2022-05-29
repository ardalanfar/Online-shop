package public_store

//Account Store intractor

import (
	"Farashop/internal/adapter/store"
	"Farashop/internal/adapter/store/model"

	"Farashop/internal/entity"
	"context"
)

func (s store.DbConn) CreateUser(ctx context.Context, user entity.User) (entity.User, error) {
	u := model.MapFromUserEntity(user)

	//cheek username and email in DB
	errcheek := s.Db.WithContext(ctx).Where("username = ? OR email = ?", u.Username, u.Email).First(&u).Error
	if errcheek == nil {
		return entity.User{}, errcheek
	}
	//create user
	errcreate := s.Db.WithContext(ctx).Create(&u).Error
	if errcreate != nil {
		return entity.User{}, errcreate
	}

	return model.MapToUserEntity(u), nil
}

func (s store.DbConn) GetUserByUsername(ctx context.Context, user entity.User) (entity.User, error) {
	u := model.MapFromUserEntity(user)

	//get id,password,username by username
	err := s.Db.WithContext(ctx).Select("id", "password", "username").Where("username = ?", u.Username).First(&u).Error
	if err != nil {
		return entity.User{}, err
	}

	return model.MapToUserEntity(u), nil
}
