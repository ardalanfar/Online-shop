package account

//Account Store intractor

import (
	"Farashop/internal/adapter/store"
	"Farashop/internal/adapter/store/model"
	"Farashop/internal/entity"
	"context"
)

func (p store.DbConn) CreatedAccount(ctx context.Context, user entity.User) (entity.User, error) {

	u := model.MapFromUserEntity(user)

	if err := p.Db.WithContext(ctx).Create(&u).Error; err != nil {
		return entity.User{}, err
	}

	return model.MapToUserEntity(u), nil
}
