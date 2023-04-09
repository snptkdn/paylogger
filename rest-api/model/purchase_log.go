package model

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type PurchaseLog struct {
	gorm.Model
	Price      int
	CategoryID uint
	Category   Category
	Date       time.Time `sql:"not null;type:date"`
}
