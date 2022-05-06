package store

import (
	"Farashop/internal/adapter/store/model"
	"Farashop/internal/conf"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//connect to the postgres

type DbConn struct {
	Db *gorm.DB
}

func New() DbConn {
	config := conf.GetConfig()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai",
		config.DB.Host,
		config.DB.Username,
		config.DB.Password,
		config.DB.Dbname,
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	if aErr := database.AutoMigrate(&model.User{}); aErr != nil {
		panic("Failed to auto migrate database!")
	}

	return DbConn{Db: database}
}
