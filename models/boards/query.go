package boards

import (
	"github.com/gin-gonic/gin"
	"github.com/pinem/server/db"
	"github.com/pinem/server/errors"
	"github.com/pinem/server/models"
	"github.com/pinem/server/utils"
	"github.com/pinem/server/utils/auth"
)

func GetAllForUser(c *gin.Context) ([]models.Board, error) {
	var boards []models.Board
	user := auth.GetUserFromContext(c)

	if err := db.ORM.Where("user_id = ?", user.ID).Find(&boards).Error; err != nil {
		return nil, errors.ErrInternalServer
	}
	return boards, nil
}

func GetOneForUser(c *gin.Context) (*models.Board, error) {
	var board models.Board
	user := auth.GetUserFromContext(c)
	boardID := utils.GetParamID(c)

	if err := db.ORM.Where("user_id = ? AND id = ?", user.ID, boardID).First(&board).Error; err != nil {
		return nil, errors.ErrNotFound
	}
	return &board, nil
}
