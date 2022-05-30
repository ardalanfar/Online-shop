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
	user = model.User{Username: "admin", Email: "admin@yahoo.com", Password: pass, ID_access: 1}
	access := []model.Access{{Access: 1, Describe: "Admin"}, {Access: 2, Describe: "Member"}}

	resultUser := Db.Create(&user).Error
	if resultUser != nil {
		return resultUser
	}
	//build ‌admin access level
	if resultAccess := Db.Create(&access).Error; resultAccess != nil {
		return resultAccess
	}
	//return
	return nil
}
