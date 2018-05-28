package database

import (
	"fmt"
	"track-api/src/helpers/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

func Connect() *gorm.DB {
	if Db != nil {
		return Db
	}
	var err error
	Db, err = gorm.Open("mysql", config.GetEnv("DATABASE_STRING"))
	Db.LogMode(true)
	fmt.Println(Db)
	if err != nil {
		panic(err)
	}
	return Db
}
