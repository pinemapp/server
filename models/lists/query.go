package lists

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/pinem/server/db"
	"github.com/pinem/server/errors"
	"github.com/pinem/server/models"
	"github.com/pinem/server/utils"
	"github.com/pinem/server/utils/auth"
)

func GetAllInBoard(c *gin.Context) ([]models.List, error) {
	var lists []models.List
	err := Scope(c).Find(&lists).Error
	if err != nil {
		return nil, errors.ErrRecordNotFound
	}
	return lists, nil
}

func GetOneInBoard(c *gin.Context) (*models.List, error) {
	var list models.List
	listID := utils.GetIntParam("list_id", c)
	err := Scope(c).Preload("Tasks").First(&list, listID).Error
	if err != nil {
		return nil, errors.ErrRecordNotFound
	}
	return &list, nil
}

func Scope(c *gin.Context) *gorm.DB {
	user := auth.GetUserFromContext(c)
	boardID := utils.GetIntParam("board_id", c)
	return db.ORM.Joins("JOIN boards ON boards.id = lists.board_id").
		Joins("JOIN board_users ON boards.id = board_users.board_id").
		Where("boards.id = ? AND board_users.user_id = ?", boardID, user.ID)
}
