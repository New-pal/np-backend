package db

import (
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

func NewConnection() *Database {
	var db *gorm.DB
	var err error
	if viper.GetString("db.type") == "postgres" {
		db, err = gorm.Open(postgres.Open(viper.GetString("db.dsn")))
	}
	if viper.GetString("db.type") == "sqlite" {
		db, err = gorm.Open(sqlite.Open(viper.GetString("db.dsn")))
	}

	if err != nil {
		panic(err)
	}

	return &Database{db}
}
