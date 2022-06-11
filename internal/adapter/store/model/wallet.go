package model

//Wallet model database

type Wallet struct {
	ID      uint `json:"id" gorm:"primary_key,serializer:json,NOT NULL"`
	Balance uint `json:"balance" gorm:"NOT NULL"`
	User_id uint `json:"user_id" gorm:"NOT NULL,UNIQUE"`
}

/*-----------------------------------------------------*/
