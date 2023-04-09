package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Category struct {
	gorm.Model
	Name string `gorm:"unique"`
}
