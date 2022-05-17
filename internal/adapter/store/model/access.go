package model

//access model database
type Access struct {
	ID       uint   `json:"id" gorm:"primary_key,serializer:json"`
	Access   uint   `json:"access"`
	Describe string `json:"describe"`
}
