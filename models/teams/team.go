package teams

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/pinem/server/db"
	"github.com/pinem/server/errors"
	"github.com/pinem/server/models"
	"github.com/pinem/server/utils"
	"github.com/pinem/server/utils/auth"
	"github.com/pinem/server/utils/messages"
	"github.com/pinem/server/utils/validators"
	"github.com/pinem/server/utils/validators/teams"
)

func Create(userID uint, f *teamvalidator.TeamForm, msg *messages.Messages) (*models.Team, error) {
	if err := validators.Validate(f, msg); err != nil {
		return nil, err
	}
	team, err := build(f, msg)
	if err != nil {
		return nil, err
	}
	team, err = create(team, userID)
	if err != nil {
		return nil, err
	}
	return team, nil
}

func Update(userID, id uint, f *teamvalidator.UpdateTeamForm, msg *messages.Messages) (*models.Team, error) {
	team, err := Find(userID, id)
	if err != nil {
		return nil, errors.ErrRecordNotFound
	}

	if err := validators.Validate(f, msg); err != nil {
		return nil, err
	}
	if err := assign(team, f, msg); err != nil {
		return nil, err
	}
	if err := db.ORM.Save(team).Error; err != nil {
		return nil, errors.GetDBError(err)
	}
	return team, nil
}

func Delete(userID, id uint) error {
	team, err := Find(userID, id)
	if err != nil {
		return err
	}
	if err := db.ORM.Delete(team).Error; err != nil {
		return errors.GetDBError(err)
	}
	return nil
}

func DeleteFromContext(c *gin.Context) error {
	user := auth.GetUserFromContext(c)
	teamID := utils.GetIntParam("team_id", c)
	return Delete(user.ID, teamID)
}

func CreateFromContext(c *gin.Context) (*models.Team, error) {
	var f teamvalidator.TeamForm
	msg := messages.GetMessages(c)
	user := auth.GetUserFromContext(c)
	c.Bind(&f)
	return Create(user.ID, &f, msg)
}

func UpdateFromContext(c *gin.Context) (*models.Team, error) {
	var f teamvalidator.UpdateTeamForm
	msg := messages.GetMessages(c)
	c.Bind(&f)

	user := auth.GetUserFromContext(c)
	teamID := utils.GetIntParam("team_id", c)
	return Update(user.ID, teamID, &f, msg)
}

func create(team *models.Team, userID uint) (*models.Team, error) {
	err := db.Transaction(db.ORM, func(tx *gorm.DB) error {
		if err := tx.Create(team).Error; err != nil {
			return errors.GetDBError(err)
		}

		teamUser := models.TeamUser{
			TeamID: team.ID,
			UserID: userID,
			Role:   models.TeamLeader,
		}
		if err := tx.Create(&teamUser).Error; err != nil {
			return errors.GetDBError(err)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return team, nil
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
