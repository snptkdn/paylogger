package util

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

func GetDb() *gorm.DB {
	db, err := gorm.Open("mysql", "snptkdn:Ryouta0820@/paylogger?parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
