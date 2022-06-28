package model

//Wallet database model

type Wallet struct {
	ID      uint `json:"id" gorm:"primary_key,serializer:json,NOT NULL"`
	Balance uint `json:"balance" gorm:"NOT NULL"`
	User_id uint `json:"user_id" gorm:"NOT NULL,UNIQUE"`
}

/*-----------------------------------------------------*/
//convert data model to database model

/*-----------------------------------------------------*/
