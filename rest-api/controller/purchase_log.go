package controller

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/snptkdn/paylogger-go/service"
)

func PostPurchaseLogHandler(c *gin.Context) {
	purchase_log, err := service.InsertPurchaseLog(c.PostForm("price"), c.PostForm("category"), c.PostForm("date"))

	if err != nil {
		c.String(500, "error occured:%s", err)
		return
	}

	c.JSON(200, purchase_log)
}

func GetPurchaseLogHandler(c *gin.Context) {
	logs, err := service.GetAllPurchaseLog()

	if err != nil {
		c.String(500, "error occured:%s", err)
		return
	}

	c.JSON(200, logs)
}
