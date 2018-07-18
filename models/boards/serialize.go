package boards

import (
	"github.com/pinem/server/models"
)

func GetSimpleBoards(boards []models.Board) []models.SimpleBoard {
	simpleBoards := []models.SimpleBoard{}
	for _, board := range boards {
		simpleBoards = append(simpleBoards, GetSimpleBoard(board))
	}
	return simpleBoards
}

func GetSimpleBoard(board models.Board) models.SimpleBoard {
	return models.SimpleBoard{Board: &board}
}
