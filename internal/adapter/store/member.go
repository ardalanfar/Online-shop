package store

import (
	"Farashop/internal/entity"
	"context"
)

func (s DbConn) ShowOrders(ctx context.Context) ([]entity.Order, error) {
	var orders []entity.Order
	//id :=

	//get all
	err := s.Db.WithContext(ctx).Select("product_id", "nimber", "buy_time", "status").Where("user_id", 2).Find(&orders)
	if err.Error != nil {
		return nil, err.Error
	}

	//return
	return orders, nil
}
