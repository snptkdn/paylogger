package service

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/snptkdn/paylogger-go/model"
	"github.com/snptkdn/paylogger-go/util"
	"time"
)

func GetTotalAmount(start_date time.Time, end_date time.Time) (int, error) {
	db := util.GetDb()

	var sum int
	db.Model(&model.PurchaseLog{}).
		Select("sum(price)").
		Where("date BETWEEN ? AND ?", start_date, end_date).
		Row().
		Scan(&sum)

	return sum, nil
}
