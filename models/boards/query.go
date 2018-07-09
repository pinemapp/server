package boards

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/pinem/server/db"
	"github.com/pinem/server/errors"
	"github.com/pinem/server/models"
	"github.com/pinem/server/utils"
	"github.com/pinem/server/utils/auth"
)

func GetAllForUser(c *gin.Context) ([]models.Board, error) {
	var boards []models.Board
	if err := scope(c).Find(&boards).Error; err != nil {
		return nil, errors.ErrInternalServer
	}
	return boards, nil
}

func GetOneForUser(c *gin.Context) (*models.Board, error) {
	var board models.Board
	boardID := utils.GetParamID(c)
	if err := scope(c).Where("boards.id = ?", boardID).First(&board).Error; err != nil {
		return nil, errors.ErrNotFound
	}
	return &board, nil
}

func scope(c *gin.Context) *gorm.DB {
	user := auth.GetUserFromContext(c)
	return db.ORM.Select("DISTINCT boards.*").Joins("LEFT JOIN board_users ON boards.id = board_users.board_id").
		Where("boards.user_id = ? OR board_users.user_id = ?", user.ID, user.ID).Order("created_at DESC")
}
