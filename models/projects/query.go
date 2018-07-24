package projects

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/pinem/server/db"
	"github.com/pinem/server/errors"
	"github.com/pinem/server/models"
	"github.com/pinem/server/utils"
	"github.com/pinem/server/utils/auth"
)

func GetAllForUser(c *gin.Context) ([]models.Project, error) {
	var projects []models.Project
	if err := Scope(c).Find(&projects).Error; err != nil {
		return nil, errors.GetDBError(err)
	}
	return projects, nil
}

func GetOneForUser(c *gin.Context) (*models.Project, error) {
	var project models.Project
	if err := getOne(c).Preload("Members").Preload("Lists", func(db *gorm.DB) *gorm.DB {
		return db.Preload("Tasks", func(db1 *gorm.DB) *gorm.DB {
			return db1.Order("tasks.order ASC")
		}).Order("lists.order ASC")
	}).First(&project).Error; err != nil {
		return nil, errors.ErrRecordNotFound
	}
	return &project, nil
}

func Scope(c *gin.Context) *gorm.DB {
	user := auth.GetUserFromContext(c)
	return db.ORM.Joins("JOIN project_users ON projects.id = project_users.project_id").
		Where("project_users.user_id = ? AND project_users.deleted_at IS NULL", user.ID).
		Order("projects.created_at DESC")
}

func getOne(c *gin.Context) *gorm.DB {
	projectID := utils.GetIntParam("project_id", c)
	return Scope(c).Where("projects.id = ?", projectID)
}
