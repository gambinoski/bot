package database

import (
	"fmt"
	"log"

	"github.com/gambinoski/bot/config"
	"github.com/jinzhu/gorm"
)

type Database struct {
	*gorm.DB
}

var (
	Db Database
)

func Connect() {
	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", config.Conf.Db.Host, config.Conf.Db.Port, config.Conf.Db.User, config.Conf.Db.Name, config.Conf.Db.Password))
	if err != nil {
		log.Fatal("failed to connect database", err)
	}
	defer db.Close()

	db.Set("gorm:table_options", "charset=utf8mb4")
	db.BlockGlobalUpdate(true)

	Db = Database{db}
}

func Setup() {
	Db.AutoMigrate(
		Guild{},
	)
}

func IsConnected(ch chan bool) {
	ch <- Db.DB.DB().Ping() == nil
}
