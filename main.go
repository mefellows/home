package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mefellows/home/config"
	"github.com/mefellows/home/controllers"
)

func main() {
	r := gin.Default()
	c := config.NewConfig()

	// Perform migrations
	//db.Migrate(c.ConnectionString)
	healthController := controllers.NewHealthController(c)
	listController := controllers.NewListController(c)

	// Routes
	r.GET("/health", healthController.Get)

	// Shopping list routes
	r.GET("/shopping/list", listController.List)
	r.GET("/shopping/list/:id/items", listController.GetItems)
	r.POST("/shopping/list/append", listController.AppendItem)
	r.PUT("/shopping/list/complete", listController.CompleteList)

	r.Run()
}
