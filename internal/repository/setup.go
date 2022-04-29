package repository

/*
import (
	"Farashop/internal/conf"
	"Farashop/internal/repository/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PsqlStore struct {
	db *gorm.DB
}

func New() PsqlStore {
	database, err := gorm.Open(postgres.Open(conf.Postgresql), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	if aErr := database.AutoMigrate(&model.User{}); aErr != nil {
		panic("Failed to auto migrate database!")
	}

	return PsqlStore{db: database}
}
*/
