package controllers

import (
	"fmt"
	. "gin-example/models"
	. "gin-example/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetBoard(context *gin.Context) {
	var board Board
	if err := DB.First(&board, context.Param("id")).Error; err != nil {
		context.AbortWithStatusJSON(http.StatusNotFound, MakeErrorMessage(err))
		return
	}
	context.JSON(http.StatusOK, board)
}

func GetBoards(context *gin.Context) {
	var boards []Board
	if err := DB.Find(&boards).Error; err != nil {
		context.AbortWithStatusJSON(http.StatusNotFound, MakeErrorMessage(err))
		return
	}
	context.JSON(http.StatusOK, boards)
}

func AddBoard(context *gin.Context) {
	var board Board
	if err := context.BindJSON(&board); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, MakeErrorMessage(err))
		return
	}
	if err := DB.Create(&board).Error; err != nil {
		context.AbortWithStatusJSON(http.StatusNotFound, MakeErrorMessage(err))
		return
	}
	context.Header("Location", "/v1/boards/"+fmt.Sprint(board.ID))
	context.JSON(http.StatusCreated, board)
}

func SetBoard(context *gin.Context) {
	var board Board
	var newBoard Board
	if err := context.BindJSON(&newBoard); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, MakeErrorMessage(err))
		return
	}
	if err := DB.First(&board, context.Param("id")).Error; err != nil {
		context.AbortWithStatusJSON(http.StatusNotFound, MakeErrorMessage(err))
		return
	}
	DB.Model(&board).Updates(newBoard)
	context.JSON(http.StatusOK, board)
}

func RemoveBoard(context *gin.Context) {
	DB.Unscoped().Delete(&Board{}, context.Param("id"))
	context.Status(http.StatusNoContent)
}
