package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Board struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

func main() {
	router := gin.Default()
	router.GET("/boards", func(context *gin.Context) {

	})
	router.POST("/boards", func(context *gin.Context) {
		var board Board
		err := context.ShouldBindJSON(&board)

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"title":   board.Title,
			"content": board.Content,
		})
	})
	router.GET("/boards/:id", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"id": context.Param("id"),
		})
	})
	router.PUT("/boards/:id", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"id": context.Param("id"),
		})
	})
	router.DELETE("/boards/:id", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"id": context.Param("id"),
		})
	})
	_ = router.Run()
}
