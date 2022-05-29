package admin

import (
	"Farashop/internal/entity"
	"context"
)

func (s DbConn) ShowMembers(ctx context.Context) ([]entity.User, error) {

	var users string

	if err := s.DB.WithContext(ctx).Find(&users).Error; err != nil {
		return nil, err
	}
}

// func (s DbConn) DeleteMember(ctx context.Context, userId uint) error {

// }
