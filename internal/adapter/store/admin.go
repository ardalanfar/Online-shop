package store

import (
	"Farashop/internal/adapter/store/model"
	"Farashop/internal/entity"
	"context"
)

func (s DbConn) ShowMembers(ctx context.Context) ([]entity.User, error) {
	var users []entity.User

	//get all id,email,username,password
	err := s.Db.WithContext(ctx).Select("id", "email", "username", "password").Find(&users)
	if err.Error != nil {
		return nil, err.Error
	}
	//return
	return users, nil
}

func (s DbConn) DeleteMember(ctx context.Context, user entity.User) error {
	u := model.MapFromUserEntity(user)

	//delete username by id
	err := s.Db.WithContext(ctx).Delete(&u)
	if err.Error != nil {
		return err.Error
	}
	//return
	return nil
}

func (s DbConn) ShowInfoMember(ctx context.Context, user entity.User) (entity.User, error) {
	u := model.MapFromUserEntity(user)

	//get "id", "email", "username" by username
	err := s.Db.WithContext(ctx).Select("id", "email", "username").First(&u)
	if err.Error != nil {
		return entity.User{}, err.Error
	}
	//return
	return model.MapToUserEntity(u), nil
}
