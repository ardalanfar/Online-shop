package store

import (
	"Farashop/internal/adapter/store/model"

	"gorm.io/gorm"
)

func InsertDefultAdmin(Db *gorm.DB) error {
	user := model.User{Username: "admin", Email: "admin@yahoo.com", Password: "123456", ID_access: 1}
	access := []model.Access{{Access: 1, Describe: "Admin"}, {Access: 2, Describe: "Member"}}
	name := ""
	//check admin

	Db.Where("username = ?", user.Username).First(&name)
	if name != "" {
		return nil
	}

	//create admin system
	resultUser := Db.Create(&user).Error
	if resultUser != nil {
		return resultUser
	}

	//build ‌admin access level
	if resultAccess := Db.Create(&access).Error; resultAccess != nil {
		return resultAccess
	}
	return nil
}

// func InsertDefultProduct(Db *gorm.DB) error {

// }
