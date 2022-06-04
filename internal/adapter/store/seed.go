package store

import (
	"Farashop/internal/adapter/store/model"
	"Farashop/pkg/encrypt"

	"gorm.io/gorm"
)

func InsertSeedAdmin(Db *gorm.DB) error {
	user := model.User{}

	//check admin
	if res := Db.Where("username = ?", "admin").Find(&user); res.RowsAffected != 0 {
		return nil
	}
	//create admin system
	pass, _ := encrypt.HashPassword("123456")
	user = model.User{Username: "admin", Email: "admin@yahoo.com", Password: pass, Access: 1}

	resultUser := Db.Create(&user).Error
	if resultUser != nil {
		return resultUser
	}

	//return
	return nil
}
