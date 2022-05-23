package model

//product model database
type Product struct {
	ID       uint   `json:"id" gorm:"primary_key,serializer:json"`
	Name     string `json:"name"`
	Price    uint   `json:"price"`
	Number   string `json:"number"`
	Describe string `json:"describe"`
}
