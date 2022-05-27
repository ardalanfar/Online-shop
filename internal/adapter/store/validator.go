package store

import (
	"Farashop/internal/adapter/store/model"
	"Farashop/internal/entity"
	"context"

	"gorm.io/gorm"
)

func (s DbConn) DoesUserExist(ctx context.Context, Username string) (bool, error) {
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
