package store

import (
	"Farashop/internal/adapter/store/model"
	"Farashop/internal/config"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//connect to the postgres

type DbConn struct {
	Db *gorm.DB
}

func New() DbConn {
	config := config.GetConfig()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai",
		config.DB.Host,
		config.DB.Username,
		config.DB.Password,
		config.DB.Dbname,
	)
	//connect db
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	//migrate db
	if aErr := database.AutoMigrate(&model.User{}); aErr != nil {
		panic("Failed to auto migrate database!")
	}

	//return connection
	return DbConn{Db: database}
}
