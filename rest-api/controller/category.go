package controller

import (
	"github.com/gin-gonic/gin"

	"fmt"
	_ "github.com/go-sql-driver/mysql"

	"github.com/snptkdn/paylogger-go/service"
)

func PostCategoryHundler(c *gin.Context) {
	name := c.PostForm("category")
	category, err := service.InsertCategory(name)
	if err != nil {
		c.String(500, fmt.Sprintf("error occured:%s", err.Error()))
	}

	c.JSON(200, category)
}

func GetCategoryHandler(c *gin.Context) {
	categories, err := service.GetCategories()
	if err != nil {
		c.String(500, fmt.Sprintf("error occured:%s", err.Error()))
	}

	c.JSON(200, categories)
}
