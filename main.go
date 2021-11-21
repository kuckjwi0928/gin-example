package main

import (
	"gin-example/controllers"
	. "gin-example/models"
	"github.com/gin-gonic/gin"
)

func main() {
	Migrate()
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.GET("/boards", controllers.GetBoards)
		v1.POST("/boards", controllers.AddBoard)
		v1.GET("/boards/:id", controllers.GetBoard)
		v1.PUT("/boards/:id", controllers.SetBoard)
		v1.DELETE("/boards/:id", controllers.RemoveBoard)
	}
	_ = router.Run()
}
