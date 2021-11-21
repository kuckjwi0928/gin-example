package repositories

import (
	"gin-example/models"
	"gorm.io/gorm"
)

type BoardRepository struct {
	DB *gorm.DB
}

func (repo *BoardRepository) Retrieve(id string) (*models.Board, error) {
	var board models.Board
	if err := repo.DB.First(&board, id).Error; err != nil {
		return nil, err
	}
	return &board, nil
}

func (repo *BoardRepository) List() []models.Board {
	var boards []models.Board
	repo.DB.Find(&boards)
	return boards
}

func (repo *BoardRepository) Add(board *models.Board) error {
	if err := repo.DB.Create(board).Error; err != nil {
		return err
	}
	return nil
}

func (repo *BoardRepository) Set(board *models.Board, newBoard *models.Board) *models.Board {
	repo.DB.Model(board).Updates(newBoard)
	return newBoard
}

func (repo *BoardRepository) Delete(id string) {
	repo.DB.Unscoped().Delete(&models.Board{}, id)
}
