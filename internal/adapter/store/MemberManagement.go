package store

import (
	"Farashop/internal/adapter/store/model"
	"Farashop/internal/entity"
	"context"
)

func (s DbConn) ShowMembers(ctx context.Context) ([]entity.User, error) {
	var users []entity.User

	//get all id,email,username,password
	if err := s.Db.WithContext(ctx).Select("id", "email", "username", "password").Find(&users).Error; err != nil {
		return nil, err
	}
	//return
	return users, nil
}

func (s DbConn) DeleteMember(ctx context.Context, user entity.User) error {
	u := model.MapFromUserEntity(user)

	//delete username by id
	errDel := s.Db.WithContext(ctx).Delete(&u).Error
	if errDel != nil {
		return errDel
	}
	//return
	return nil
}

func (s DbConn) ShowInfoMember(ctx context.Context, user entity.User) (entity.User, error) {
	u := model.MapFromUserEntity(user)

	//get "id", "email", "username" by username
	if err := s.Db.WithContext(ctx).Select("id", "email", "username").First(&u).Error; err != nil {
		return entity.User{}, err
	}
	//return
	return model.MapToUserEntity(u), nil
}
