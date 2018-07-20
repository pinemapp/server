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
	team, err := Build(f, msg)
	if err != nil {
		return nil, err
	}
	if err := db.ORM.Create(team).Error; err != nil {
		return nil, errors.ErrInternalServer
	}
	return team, nil
}

func CreateFromContext(c *gin.Context, msg *messages.Messages) (*models.Team, error) {
	var f teamvalidator.TeamForm
	c.Bind(&f)
	return Create(&f, msg)
}

func UpdateFromContext(c *gin.Context, msg *messages.Messages) (*models.Team, error) {
	var f teamvalidator.UpdateTeamForm
	c.Bind(&f)
	teamID := utils.GetIntParam("team_id", c)
	return Update(teamID, &f, msg)
}

func Update(id uint, f *teamvalidator.UpdateTeamForm, msg *messages.Messages) (*models.Team, error) {
	var team models.Team
	if err := db.ORM.Where("id = ?", id).First(&team).Error; err != nil {
		return nil, errors.ErrRecordNotFound
	}

	if err := validators.Validate(f, msg); err != nil {
		return nil, err
	}
	if err := Assign(&team, f, msg); err != nil {
		return nil, err
	}
	if err := db.ORM.Create(&team).Error; err != nil {
		return nil, errors.ErrInternalServer
	}
	return &team, nil
}

func Build(f *teamvalidator.TeamForm, msg *messages.Messages) (*models.Team, error) {
	if err := validators.Validate(f, msg); err != nil {
		return nil, err
	}

	var team models.Team
	validators.Bind(&team, f)
	if team.Slug == "" {
		team.Slug = generateSlug(team.Name)
	}
	return &team, nil
}

func Assign(team *models.Team, f *teamvalidator.UpdateTeamForm, msg *messages.Messages) error {
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
	count := 1
	slug := utils.GenerateSlug(name)
	for isSlugExist(slug) {
		slug = fmt.Sprintf("%s%d", slug, count)
		count += 1
	}
	return slug
}

func isSlugExist(slug string) bool {
	var count int
	db.ORM.Table("teams").Where("slug = ?", slug).Count(&count)
	return count > 0
}
