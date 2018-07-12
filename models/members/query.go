package members

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/pinem/server/db"
	"github.com/pinem/server/errors"
	"github.com/pinem/server/models"
	"github.com/pinem/server/utils"
)

func GetAllInBoard(c *gin.Context) ([]models.BoardUser, error) {
	var members []models.BoardUser
	if err := Scope(c).Find(&members).Error; err != nil {
		return nil, errors.ErrInternalServer
	}
	return members, nil
}

func GetOneInBoard(c *gin.Context) (*models.BoardUser, error) {
	var member models.BoardUser
	memberID := utils.GetIntParam("member_id", c)
	if err := Scope(c).Where("board_users.id = ?", memberID).First(&member).Error; err != nil {
		return nil, errors.ErrNotFound
	}
	return &member, nil
}

func Scope(c *gin.Context) *gorm.DB {
	boardID := utils.GetIntParam("board_id", c)
	return db.ORM.Joins("JOIN boards ON boards.id = board_users.board_id").Where("board_id = ?", boardID)
}
