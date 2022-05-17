package store

import (
	"Farashop/internal/adapter/store/model"

	"gorm.io/gorm"
)

func InsertAdmin(Db *gorm.DB) error {
	user := model.User{Username: "admin", Email: "admin@yahoo.com", Password: "123456", ID_access: 1}
	access := []model.Access{{Access: 1, Describe: "Admin"}, {Access: 2, Describe: "Member"}}

	if resultUser := Db.Create(&user).Error; resultUser != nil {
		return resultUser
	}

	if resultAccess := Db.Create(&access).Error; resultAccess != nil {
		return resultAccess
	}
	return nil
}
