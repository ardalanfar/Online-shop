package model

import "time"

//Order model database

type Order struct {
	ID         uint      `json:"id" gorm:"primary_key,serializer:json,NOT NULL"`
	User_id    uint      `json:"user_id" gorm:"NOT NULL"`
	Product_id uint      `json:"product_id" gorm:"NOT NULL"`
	Buy_time   time.Time `json:"Buy_time" gorm:"NOT NULL"`
}

/*-----------------------------------------------------*/
