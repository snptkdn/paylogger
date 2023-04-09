package service

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/snptkdn/paylogger-go/model"
	"github.com/snptkdn/paylogger-go/util"
)

func InsertCategory(name string) (model.Category, error) {
	category := model.Category{
		Name: name,
	}

	db := util.GetDb()

	db.Create(&category)

	return category, nil
}

func GetCategories() ([]model.Category, error) {
	db := util.GetDb()

	var categories []model.Category
	err := db.Find(&categories).Error
	if err != nil {
		return nil, err
	}

	return categories, nil
}
