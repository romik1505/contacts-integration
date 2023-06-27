package store

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"week3_docker/internal/config"
	"week3_docker/internal/model"
)

type Store struct {
	DB *gorm.DB
}

func NewStore() (Store, error) {
	dbConf := config.Config.DBConfig
	conString := dbConf.ConnectionString()

	db, err := gorm.Open(mysql.Open(conString), &gorm.Config{})
	if err != nil {
		return Store{}, err
	}

	db.AutoMigrate(&model.Account{}, &model.Contact{}, &model.Integration{})

	return Store{
		DB: db,
	}, nil
}
