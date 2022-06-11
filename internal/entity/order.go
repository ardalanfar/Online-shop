package entity

import "time"

//Entity Order

type Order struct {
	ID         uint      `json:"id"`
	User_id    uint      `json:"user_id"`
	Product_id uint      `json:"product_id"`
	Number     uint      `json:"number"`
	Buy_time   time.Time `json:"buy_time"`
	Status     string    `json:"status"`
}
