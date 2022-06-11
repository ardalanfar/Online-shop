package store

//Account Store intractor

import (
	"Farashop/internal/adapter/store/model"

	"Farashop/internal/entity"
	"context"
)

func (s DbConn) Register(ctx context.Context, user entity.User) error {
	u := model.MapFromUserEntity(user)

	//cheek username and email
	Cheek := s.Db.WithContext(ctx).Select("id").Where("username = ? OR email = ?", u.Username, u.Email).First(&u)
	if Cheek.Error != nil && Cheek.RowsAffected != 0 && u.ID != 0 {
		return Cheek.Error
	}
	//create user
	Create := s.Db.WithContext(ctx).Create(&u)
	if Create.Error != nil {
		return Create.Error
	}
	//return
	return nil
}

func (s DbConn) Login(ctx context.Context, user entity.User) (entity.User, error) {
	u := model.MapFromUserEntity(user)

	//get "id", "email", "password", "access", "username" by username
	Cheek := s.Db.WithContext(ctx).Select("id", "email", "password", "access", "username").Where("username = ?", u.Username).First(&u)
	if Cheek.Error != nil {
		return entity.User{}, Cheek.Error
	}
	//return
	return model.MapToUserEntity(u), nil
}

func (s DbConn) MemberValidation(ctx context.Context, user entity.User) (bool, error) {
	u := model.MapFromUserEntity(user)

	//update verify code
	Cheek := s.Db.WithContext(ctx).Model(&u).Where("username = ? AND verification_code = ?", u.Username, u.Verification_code).Update("is_verified", "active")
	if Cheek.RowsAffected == 0 || Cheek.Error != nil {
		return false, Cheek.Error
	}
	//return
	return true, nil
}
