package admin_store

import (
	"Farashop/internal/adapter/store"
	"Farashop/internal/entity"
	"context"
)

func (s store.DbConn) ShowMembers(ctx context.Context) ([]entity.User, error) {
	var users []entity.User

	if err := s.Db.WithContext(ctx).Find(&users).Error; err != nil {
		return nil, err
	}
}

// func (s DbConn) DeleteMember(ctx context.Context, userId uint) error {

// }
