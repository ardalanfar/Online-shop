package store

import (
	"Farashop/internal/adapter/store/model"
	"Farashop/internal/entity"
	"context"

	"gorm.io/gorm"
)

//Does Username Exist in Database
func (s DbConn) DoesUsernameExist(ctx context.Context, Username string) (bool, error) {
	u := model.MapFromUserEntity(entity.User{})

	err := s.Db.WithContext(ctx).Where("username = ?", Username).First(&u).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

//Does Id Exist in Database
func (s DbConn) DoesIDExist(ctx context.Context, Id uint) (bool, error) {
	u := model.MapFromUserEntity(entity.User{})

	err := s.Db.WithContext(ctx).Where("id = ?", Id).First(&u).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
