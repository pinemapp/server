package members

import (
	"github.com/gin-gonic/gin"
	"github.com/pinem/server/db"
	"github.com/pinem/server/errors"
	"github.com/pinem/server/models"
	"github.com/pinem/server/utils"
	"github.com/pinem/server/utils/messages"
	"github.com/pinem/server/utils/validators"
	"github.com/pinem/server/utils/validators/members"
)

func Add(c *gin.Context, msg *messages.Messages) (*models.ProjectUser, error) {
	var f membervalidator.MemberForm
	c.Bind(&f)

	if err := validators.Validate(&f, msg); err != nil {
		return nil, err
	}

	ok, err := isMember(f.UserID, c)
	if err != nil {
		return nil, errors.ErrRecordNotFound
	} else if ok {
		return nil, errors.ErrMemberAlreadyInProject
	}

	projectID := utils.GetIntParam("project_id", c)
	member := models.ProjectUser{
		UserID:    f.UserID,
		ProjectID: projectID,
		Role:      f.Role,
	}
	if err := db.ORM.Create(&member).Error; err != nil {
		return nil, errors.ErrUserNotExist
	}

	return &member, nil
}

func Update(c *gin.Context, msg *messages.Messages) (*models.ProjectUser, error) {
	var f membervalidator.UpdateMemberForm
	c.Bind(&f)

	if err := validators.Validate(&f, msg); err != nil {
		return nil, err
	}

	member, err := GetOneInProject(c)
	if err != nil {
		return nil, err
	}

	member.Role = f.Role
	if err := db.ORM.Save(member).Error; err != nil {
		return nil, errors.GetDBError(err)
	}

	return member, nil
}

func Delete(c *gin.Context) error {
	member, err := GetOneInProject(c)
	if err != nil {
		return err
	}

	if err := db.ORM.Delete(member).Error; err != nil {
		return errors.GetDBError(err)
	}
	return nil
}

func isMember(userID uint, c *gin.Context) (bool, error) {
	var count int
	if err := Scope(c).Model(&models.ProjectUser{}).Where("project_users.user_id = ?", userID).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
