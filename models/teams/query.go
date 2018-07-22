package teams

import (
	"github.com/gin-gonic/gin"
	"github.com/pinem/server/db"
	"github.com/pinem/server/errors"
	"github.com/pinem/server/models"
	"github.com/pinem/server/utils"
)

func Find(id uint) (*models.Team, error) {
	var team models.Team
	if err := db.ORM.First(&team, id).Error; err != nil {
		return nil, errors.ErrRecordNotFound
	}
	return &team, nil
}

func FindAll() ([]models.Team, error) {
	var teams []models.Team
	if err := db.ORM.Order("created_at ASC").Find(&teams).Error; err != nil {
		return nil, errors.GetDBError(err)
	}
	return teams, nil
}

func FindFromContext(c *gin.Context) (*models.Team, error) {
	teamID := utils.GetIntParam("team_id", c)
	return Find(teamID)
}
