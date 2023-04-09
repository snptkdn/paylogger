package controller

import (
	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"

	"github.com/snptkdn/paylogger-go/model"
	"github.com/snptkdn/paylogger-go/util"
)

func GetMigrateHandler(c *gin.Context) {
	db := util.GetDb()

	db.AutoMigrate(&model.Category{})
	db.AutoMigrate(&model.PurchaseLog{}).AddForeignKey("category_id", "categories(id)", "RESTRICT", "RESTRICT")
}
