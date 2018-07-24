package teams

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/pinem/server/db"
	"github.com/pinem/server/errors"
	"github.com/pinem/server/models"
	"github.com/pinem/server/utils"
	"github.com/pinem/server/utils/auth"
)

func Find(userID, id uint) (*models.Team, error) {
	var team models.Team
	if err := scope(userID).First(&team, id).Error; err != nil {
		return nil, errors.ErrRecordNotFound
	}
	return &team, nil
}

func FindAll(userID uint) ([]models.Team, error) {
	var teams []models.Team
	if err := scope(userID).Find(&teams).Error; err != nil {
		return nil, errors.GetDBError(err)
	}
	return teams, nil
}

func FindAllFromContext(c *gin.Context) ([]models.Team, error) {
	user := auth.GetUserFromContext(c)
	return FindAll(user.ID)
}

func FindFromContext(c *gin.Context) (*models.Team, error) {
	user := auth.GetUserFromContext(c)
	teamID := utils.GetIntParam("team_id", c)
	return Find(user.ID, teamID)
}

func scope(userID uint) *gorm.DB {
	return db.ORM.Joins("JOIN team_users ON teams.id = team_users.team_id").
		Where("team_users.user_id = ?", userID).Order("created_at ASC")
}
