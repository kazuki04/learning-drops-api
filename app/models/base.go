package models

import (
	"github.com/learning-drops-api/config"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	_ "github.com/go-sql-driver/mysql"
)

var DbConnection *gorm.DB

func init() {
	var err error
	dsn := config.Config.Dsn
	DbConnection, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
}

func DBmigrate() *gorm.DB {
	DbConnection.AutoMigrate(&Section{})
	return DbConnection
}
