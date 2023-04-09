package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/snptkdn/paylogger-go/service"
	"github.com/snptkdn/paylogger-go/validator"
)

func TotalHundler(c *gin.Context) {
	year := c.Query("year")
	month := c.Query("month")
	day := c.Query("day")

	start_date, end_date, err := validator.ValidateTotal(year, month, day)
	if err != nil {
		c.String(500, "error occured:%s", err)
		return
	}

	sum, err := service.GetTotalAmount(*start_date, *end_date)
	if err != nil {
		c.String(500, "error occured:%s", err)
		return
	}

	c.String(200, strconv.Itoa(sum))
}
