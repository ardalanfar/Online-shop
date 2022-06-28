package model

//Product database model

type Product struct {
	ID       uint   `json:"id" gorm:"primary_key,serializer:json,NOT NULL"`
	Name     string `json:"name" gorm:"NOT NULL,UNIQUE"`
	Price    uint   `json:"price" gorm:"NOT NULL"`
	Number   uint   `json:"number" gorm:"NOT NULL"`
	Describe string `json:"describe"`
}

/*-----------------------------------------------------*/
//convert data model to database model

/*-----------------------------------------------------*/
