package controllers

import (
	"fmt"
	. "gin-example/models"
	"gin-example/repositories"
	. "gin-example/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BoardController struct {
	Repository *repositories.BoardRepository
}

func (api *BoardController) GetBoard(context *gin.Context) {
	board, err := api.Repository.Retrieve(context.Param("id"))
	if err != nil {
		context.AbortWithStatusJSON(http.StatusNotFound, MakeErrorMessage(err))
		return
	}
	context.JSON(http.StatusOK, board)
}

func (api *BoardController) GetBoards(context *gin.Context) {
	boards := api.Repository.List()
	if len(boards) < 1 {
		context.AbortWithStatus(http.StatusNotFound)
		return
	}
	context.JSON(http.StatusOK, boards)
}

func (api *BoardController) AddBoard(context *gin.Context) {
	var board Board
	var err error
	if err = context.BindJSON(&board); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, MakeErrorMessage(err))
		return
	}
	if err = api.Repository.Add(&board); err != nil {
		context.AbortWithStatusJSON(http.StatusNotFound, MakeErrorMessage(err))
		return
	}
	context.Header("Location", "/v1/boards/"+fmt.Sprint(board.ID))
	context.JSON(http.StatusCreated, board)
}

func (api *BoardController) SetBoard(context *gin.Context) {
	var newBoard Board
	if err := context.BindJSON(&newBoard); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, MakeErrorMessage(err))
		return
	}
	board, err := api.Repository.Retrieve(context.Param("id"))
	if err != nil {
		context.AbortWithStatusJSON(http.StatusNotFound, MakeErrorMessage(err))
		return
	}
	api.Repository.Set(board, &newBoard)
	context.JSON(http.StatusOK, board)
}

func (api *BoardController) DeleteBoard(context *gin.Context) {
	api.Repository.Delete(context.Param("id"))
	context.Status(http.StatusNoContent)
}
