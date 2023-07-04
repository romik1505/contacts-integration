package store

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
	"week3_docker/internal/config"
	"week3_docker/internal/model"
)

type Store struct {
	DB *gorm.DB
}

func NewStore() Store {
	dbConf := config.Config.DBConfig
	conString := dbConf.ConnectionString()

	var db *gorm.DB
	var err error

	for tries := 0; tries < 4; tries++ {
		db, err = gorm.Open(mysql.Open(conString), &gorm.Config{})
		if err != nil {
			log.Printf("Try[%d] connect to db: failed %v", tries+1, err)
		} else {
			break
		}
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		log.Fatalf("db connection failed")
	}
	log.Printf("db connection success")

	if err = db.AutoMigrate(&model.Account{}, &model.Contact{}, &model.Integration{}); err != nil {
		log.Fatalf("db migration failed: %v", err)
	}

	return Store{
		DB: db,
	}
}
