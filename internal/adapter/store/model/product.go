package model

//Product model database

type Product struct {
	ID       uint   `json:"id" gorm:"primary_key,serializer:json,NOT NULL"`
	Name     string `json:"name" gorm:"NOT NULL,UNIQUE"`
	Price    uint   `json:"price" gorm:"NOT NULL"`
	Number   uint   `json:"number" gorm:"NOT NULL"`
	Describe string `json:"describe"`
}

/*-----------------------------------------------------*/
/*-----------------------------------------------------*/
