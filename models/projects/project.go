package projects

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/pinem/server/db"
	"github.com/pinem/server/errors"
	"github.com/pinem/server/models"
	"github.com/pinem/server/utils/auth"
	"github.com/pinem/server/utils/messages"
	"github.com/pinem/server/utils/validators"
	"github.com/pinem/server/utils/validators/projects"
)

func Create(c *gin.Context, msg *messages.Messages) (*models.Project, error) {
	var f projectvalidator.ProjectForm
	c.Bind(&f)
	if err := validators.Validate(&f, msg); err != nil {
		return nil, err
	}

	project, err := create(&f, c)
	if err != nil {
		return nil, errors.GetDBError(err)
	}
	return project, nil
}

func Update(c *gin.Context, msg *messages.Messages) (*models.Project, error) {
	project, err := GetOneForUser(c)
	if err != nil {
		return nil, err
	}

	var f projectvalidator.UpdateProjectForm
	c.Bind(&f)
	if err := validators.Validate(&f, msg); err != nil {
		return nil, err
	}

	validators.Bind(project, &f)
	if err := db.ORM.Save(project).Error; err != nil {
		return nil, errors.GetDBError(err)
	}
	return project, nil
}

func Delete(c *gin.Context) error {
	project, err := GetOneForUser(c)
	if err != nil {
		return err
	}

	err = db.Transaction(db.ORM, func(tx *gorm.DB) error {
		if err := tx.Delete(models.Task{}, "project_id = ?", project.ID).Error; err != nil {
			return err
		}
		if err := tx.Delete(models.List{}, "project_id = ?", project.ID).Error; err != nil {
			return err
		}
		if err := tx.Delete(models.ProjectUser{}, "project_id = ?", project.ID).Error; err != nil {
			return err
		}
		if err := tx.Delete(&project).Error; err != nil {
			return err
		}
		return nil
	})
	return errors.GetDBError(err)
}

func create(f *projectvalidator.ProjectForm, c *gin.Context) (*models.Project, error) {
	user := auth.GetUserFromContext(c)
	var project models.Project
	validators.Bind(&project, f)

	err := db.Transaction(db.ORM, func(tx *gorm.DB) error {
		if err := tx.Create(&project).Error; err != nil {
			return errors.GetDBError(err)
		}
		projectMember := models.ProjectUser{
			UserID:    user.ID,
			ProjectID: project.ID,
			Role:      models.ProjectOwner,
		}
		if err := tx.Create(&projectMember).Error; err != nil {
			return errors.GetDBError(err)
		}
		project.Members = append(project.Members, projectMember)
		return nil
	})

	if err != nil {
		return nil, err
	}
	return &project, nil
}
