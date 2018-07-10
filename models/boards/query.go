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
	if err := Scope(c).Preload("Members").Find(&boards).Error; err != nil {
		return nil, errors.ErrInternalServer
	}
	return boards, nil
}

func GetOneForUser(c *gin.Context) (*models.Board, error) {
	var board models.Board
	if err := getOne(c).Preload("Members").First(&board).Error; err != nil {
		return nil, errors.ErrNotFound
	}
	return &board, nil
}

func Scope(c *gin.Context) *gorm.DB {
	user := auth.GetUserFromContext(c)
	return db.ORM.Select("DISTINCT boards.*").Joins("LEFT JOIN board_users ON boards.id = board_users.board_id").
		Where("boards.user_id = ? OR board_users.user_id = ?", user.ID, user.ID).Order("created_at DESC")
}

func getOne(c *gin.Context) *gorm.DB {
	boardID := utils.GetIntParam("board_id", c)
	return Scope(c).Where("boards.id = ?", boardID)
}
