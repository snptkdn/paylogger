package validator

import (
	"github.com/snptkdn/paylogger-go/model"
	"strconv"
	"time"
)

func ValidatePurchaseLog(raw_price string, raw_category string, raw_date string) (*model.PurchaseLog, error) {
	price, err := strconv.Atoi(raw_price)
	if err != nil {
		return nil, err
	}

	category, err := strconv.ParseUint(raw_category, 10, 64)
	if err != nil {
		return nil, err
	}

	date, err := time.Parse("20060102", raw_date)
	if err != nil {
		return nil, err
	}

	purchase_log := model.PurchaseLog{
		Price:      price,
		CategoryID: uint(category),
		Date:       date,
	}

	return &purchase_log, nil
}
