package tasks

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/pinem/server/db"
	"github.com/pinem/server/errors"
	"github.com/pinem/server/models"
	"github.com/pinem/server/utils"
)

func GetAllInProject(c *gin.Context) ([]models.Task, error) {
	var tasks []models.Task
	err := Scope(c).Find(&tasks).Error
	if err != nil {
		return nil, errors.GetDBError(err)
	}
	return tasks, nil
}

func GetOneInProject(c *gin.Context) (*models.Task, error) {
	var task models.Task
	taskID := utils.GetIntParam("task_id", c)
	err := Scope(c).Where("tasks.id = ?", taskID).First(&task).Error
	if err != nil {
		return nil, errors.GetDBError(err)
	}
	return &task, nil
}

func Scope(c *gin.Context) *gorm.DB {
	projectID := utils.GetIntParam("project_id", c)
	return db.ORM.Joins("JOIN projects ON projects.id = tasks.project_id").
		Where("tasks.project_id = ?", projectID)
}
