package store

import (
	"Farashop/internal/adapter/store/model"
	"Farashop/internal/entity"
	"context"
)

func (s DbConn) ShowOrders(ctx context.Context, userID uint) ([]entity.Order, error) {
	var orders []model.Order

	//get all
	cheek := s.Db.WithContext(ctx).Table("orders").Select("orders.id", "orders.number", "orders.buy_time", "orders.status", "products.name").Joins("JOIN products ON products.id = orders.product_id").Where("user_id", userID).Find(&orders)
	orderEntities := make([]entity.Order, len(orders))

	if cheek.Error != nil {
		return nil, cheek.Error
	}

	for i := range orders {
		orderEntities[i] = model.MapToOrderEntity(orders[i])
	}

	//return
	return orderEntities, nil
}
