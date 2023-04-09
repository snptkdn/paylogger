package service

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/snptkdn/paylogger-go/model"
	"github.com/snptkdn/paylogger-go/util"
	"github.com/snptkdn/paylogger-go/validator"
)

func InsertPurchaseLog(price string, category string, date string) (*model.PurchaseLog, error) {
	purchase_log, err := validator.ValidatePurchaseLog(price, category, date)
	if err != nil {
		return nil, err
	}

	db := util.GetDb()
	db.Create(&purchase_log)

	return purchase_log, nil
}

func GetAllPurchaseLog() ([]model.PurchaseLog, error) {
	db := util.GetDb()

	var logs []model.PurchaseLog
	err := db.Find(&logs).Error
	if err != nil {
		return nil, err
	}

	return logs, nil
}
