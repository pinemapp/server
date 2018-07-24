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

func GetAllInProject(c *gin.Context) ([]models.List, error) {
	var lists []models.List
	err := Scope(c).Find(&lists).Error
	if err != nil {
		return nil, errors.ErrRecordNotFound
	}
	return lists, nil
}

func GetOneInProject(c *gin.Context) (*models.List, error) {
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
	projectID := utils.GetIntParam("project_id", c)
	return db.ORM.Joins("JOIN projects ON projects.id = lists.project_id").
		Joins("JOIN project_users ON projects.id = project_users.project_id").
		Where("projects.id = ? AND project_users.user_id = ?", projectID, user.ID)
}
