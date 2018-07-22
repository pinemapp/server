package teams

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pinem/server/db"
	"github.com/pinem/server/errors"
	"github.com/pinem/server/models"
	"github.com/pinem/server/utils"
	"github.com/pinem/server/utils/messages"
	"github.com/pinem/server/utils/validators"
	"github.com/pinem/server/utils/validators/teams"
)

func Create(f *teamvalidator.TeamForm, msg *messages.Messages) (*models.Team, error) {
	if err := validators.Validate(f, msg); err != nil {
		return nil, err
	}
	team, err := build(f, msg)
	if err != nil {
		return nil, err
	}
	if err := db.ORM.Create(team).Error; err != nil {
		return nil, errors.ErrInternalServer
	}
	return team, nil
}

func Update(id uint, f *teamvalidator.UpdateTeamForm, msg *messages.Messages) (*models.Team, error) {
	var team models.Team
	if err := db.ORM.Where("id = ?", id).First(&team).Error; err != nil {
		return nil, errors.ErrRecordNotFound
	}

	if err := validators.Validate(f, msg); err != nil {
		return nil, err
	}
	if err := assign(&team, f, msg); err != nil {
		return nil, err
	}
	if err := db.ORM.Save(&team).Error; err != nil {
		return nil, errors.GetDBError(err)
	}
	return &team, nil
}

func Delete(id uint) error {
	team, err := Find(id)
	if err != nil {
		return err
	}
	if err := db.ORM.Delete(team).Error; err != nil {
		return errors.GetDBError(err)
	}
	return nil
}

func DeleteFromContext(c *gin.Context) error {
	teamID := utils.GetIntParam("team_id", c)
	return Delete(teamID)
}

func CreateFromContext(c *gin.Context) (*models.Team, error) {
	var f teamvalidator.TeamForm
	msg := messages.GetMessages(c)
	c.Bind(&f)
	return Create(&f, msg)
}

func UpdateFromContext(c *gin.Context) (*models.Team, error) {
	var f teamvalidator.UpdateTeamForm
	msg := messages.GetMessages(c)
	c.Bind(&f)
	teamID := utils.GetIntParam("team_id", c)
	return Update(teamID, &f, msg)
}

func build(f *teamvalidator.TeamForm, msg *messages.Messages) (*models.Team, error) {
	if err := validators.Validate(f, msg); err != nil {
		return nil, err
	}

	var team models.Team
	validators.Bind(&team, f)
	if f.Slug == nil {
		team.Slug = generateSlug(team.Name)
	}
	return &team, nil
}

func assign(team *models.Team, f *teamvalidator.UpdateTeamForm, msg *messages.Messages) error {
	if err := validators.Validate(f, msg); err != nil {
		return err
	}

	validators.Bind(team, f)
	if f.Slug != nil {
		if isSlugExist(*f.Slug) {
			msg.ErrorT("slug", errors.ErrNotUnique)
			return errors.ErrNotUnique
		}
	}
	return nil
}

func generateSlug(name string) string {
	slug := utils.GenerateSlug(name)
	for isSlugExist(slug) {
		slug = fmt.Sprintf("%s%d", slug, utils.RandomNumString(8))
	}
	return slug
}

func isSlugExist(slug string) bool {
	var count int
	db.ORM.Table("teams").Where("slug = ?", slug).Count(&count)
	return count > 0
}
