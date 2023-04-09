package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/appengine"

	"github.com/snptkdn/paylogger-go/controller"
)

func main() {
	engine := gin.Default()
	engine.GET("/", controller.IndexHandler)
	engine.GET("/migrate", controller.GetMigrateHandler)
	engine.POST("/purchase_log", controller.PostPurchaseLogHandler)
	engine.GET("/purchase_log", controller.GetPurchaseLogHandler)
	engine.GET("/category", controller.GetCategoryHandler)
	engine.POST("/category", controller.PostCategoryHundler)
	engine.GET("/total", controller.TotalHundler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	engine.Run()
	http.Handle("/", engine)
	appengine.Main()
}
