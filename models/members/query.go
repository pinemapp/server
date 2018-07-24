package members

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/pinem/server/db"
	"github.com/pinem/server/errors"
	"github.com/pinem/server/models"
	"github.com/pinem/server/utils"
)

func GetAllInProject(c *gin.Context) ([]models.ProjectUser, error) {
	var members []models.ProjectUser
	if err := Scope(c).Find(&members).Error; err != nil {
		return nil, errors.GetDBError(err)
	}
	return members, nil
}

func GetOneInProject(c *gin.Context) (*models.ProjectUser, error) {
	var member models.ProjectUser
	memberID := utils.GetIntParam("member_id", c)
	if err := Scope(c).Where("project_users.id = ?", memberID).First(&member).Error; err != nil {
		return nil, errors.ErrRecordNotFound
	}
	return &member, nil
}

func Scope(c *gin.Context) *gorm.DB {
	projectID := utils.GetIntParam("project_id", c)
	return db.ORM.Joins("JOIN projects ON projects.id = project_users.project_id").Where("project_id = ?", projectID)
}
