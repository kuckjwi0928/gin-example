package main

import (
	"gin-example/controllers"
	"gin-example/models"
	"gin-example/repositories"
	"github.com/gin-gonic/gin"
)

func main() {
	models.Init()
	router := gin.Default()

	boardApi := controllers.BoardController{
		Repository: &repositories.BoardRepository{DB: models.GetDB()},
	}

	v1 := router.Group("/v1")
	{
		v1.GET("/boards", boardApi.GetBoards)
		v1.POST("/boards", boardApi.AddBoard)
		v1.GET("/boards/:id", boardApi.GetBoard)
		v1.PUT("/boards/:id", boardApi.SetBoard)
		v1.DELETE("/boards/:id", boardApi.DeleteBoard)
	}
	_ = router.Run()
}
