package repositories

import (
	"gin-example/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"testing"
)

var mockRepository BoardRepository

func init() {
	gin.SetMode(gin.TestMode)
	models.TestInit()
	board := models.Board{Title: "kuckjwi", Content: "Hi Kuckjwi!"}
	models.GetDB().Create(&board)
	mockRepository = BoardRepository{DB: models.GetDB()}
}

func TestBoardRepository_Add(t *testing.T) {
	board := models.Board{Title: "isis", Content: "Hi isis!"}
	_ = mockRepository.Add(&board)
	assert.NotNil(t, board.ID)
	assert.NotNil(t, board.CreatedAt)
}

func TestBoardRepository_Set(t *testing.T) {
	board, _ := mockRepository.Retrieve("1")
	newBoard := models.Board{Title: "what are you ding?", Content: "umm..."}
	mockRepository.Set(board, &newBoard)
	assert.Equal(t, board.Title, "what are you ding?")
	assert.Equal(t, board.Content, "umm...")
}

func TestBoardRepository_Delete(t *testing.T) {
	mockRepository.Delete("1")
	board, _ := mockRepository.Retrieve("1")
	assert.Nil(t, board)
}

func TestBoardRepository_Retrieve(t *testing.T) {
	board, _ := mockRepository.Retrieve("1")
	assert.Equal(t, board.Title, "kuckjwi")
	assert.Equal(t, board.Content, "Hi Kuckjwi!")
}

func TestBoardRepository_List(t *testing.T) {
	assert.NotEmpty(t, mockRepository.List())
}
