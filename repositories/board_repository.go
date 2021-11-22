package repositories

import (
	"gin-example/models"
	"gorm.io/gorm"
)

type BoardRepository struct {
	DB *gorm.DB
}

func (b *BoardRepository) Retrieve(id string) (*models.Board, error) {
	var board models.Board
	if err := b.DB.First(&board, id).Error; err != nil {
		return nil, err
	}
	return &board, nil
}

func (b *BoardRepository) List() []models.Board {
	var boards []models.Board
	b.DB.Find(&boards)
	return boards
}

func (b *BoardRepository) Add(board *models.Board) error {
	if err := b.DB.Create(board).Error; err != nil {
		return err
	}
	return nil
}

func (b *BoardRepository) Set(board *models.Board, newBoard *models.Board) {
	b.DB.Model(board).Updates(newBoard)
}

func (b *BoardRepository) Delete(id string) {
	b.DB.Unscoped().Delete(&models.Board{}, id)
}
