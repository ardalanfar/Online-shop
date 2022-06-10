package store

import (
	"Farashop/internal/entity"
	"context"
)

//Does Username Exist and active in Database
func (conn DbConn) DoesUsernameActiveStore(ctx context.Context, Username string) (bool, error) {
	cheek := conn.Db.WithContext(ctx).Where("username = ? AND is_verified = ?", Username, "active").First(&entity.User{})
	if cheek.Error != nil || cheek.RowsAffected == 0 {
		return false, cheek.Error
	}
	return true, nil
}

//Does Id Exist in Database
func (conn DbConn) DoesIDExistStore(ctx context.Context, Id uint) (bool, error) {
	cheek := conn.Db.WithContext(ctx).Where("id = ?", Id).First(&entity.User{})
	if cheek.Error != nil || cheek.RowsAffected == 0 {
		return false, cheek.Error
	}
	return true, nil
}
