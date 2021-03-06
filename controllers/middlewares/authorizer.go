package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/pinem/server/controllers/router"
	"github.com/pinem/server/db"
	"github.com/pinem/server/models"
	"github.com/pinem/server/utils"
	"github.com/pinem/server/utils/auth"
)

func Authorizer() gin.HandlerFunc {
	e := router.GetEnforcer()

	return func(c *gin.Context) {
		user := auth.GetUserFromContext(c)
		if user == nil {
			router.RenderForbidden(c)
			return
		}

		path := c.Request.URL.Path
		method := c.Request.Method
		projectID := utils.GetIntParam("project_id", c)

		if projectID == 0 {
			ok, err := e.EnforceSafe("user", path, method)
			if err != nil || !ok {
				router.RenderForbidden(c)
				return
			}
			c.Next()
			return
		}

		member, err := getMember(user.ID, projectID)
		if err != nil {
			router.RenderForbidden(c)
			return
		}

		ok, err := e.EnforceSafe(string(member.Role), path, method)
		if err != nil || !ok {
			router.RenderForbidden(c)
			return
		}

		c.Next()
	}
}

func TeamAuthorizer() gin.HandlerFunc {
	e := router.GetEnforcer()

	return func(c *gin.Context) {
		user := auth.GetUserFromContext(c)
		if user == nil {
			router.RenderForbidden(c)
			return
		}

		path := c.Request.URL.Path
		method := c.Request.Method
		teamID := utils.GetIntParam("team_id", c)

		if teamID == 0 {
			ok, err := e.EnforceSafe("user", path, method)
			if err != nil || !ok {
				router.RenderForbidden(c)
				return
			}
			c.Next()
			return
		}

		member, err := getTeamMember(user.ID, teamID)
		if err != nil {
			router.RenderForbidden(c)
			return
		}

		ok, err := e.EnforceSafe(string(member.Role), path, method)
		if err != nil || !ok {
			router.RenderForbidden(c)
			return
		}

		c.Next()
	}
}

func getMember(userID, projectID uint) (*models.ProjectUser, error) {
	var member models.ProjectUser
	err := db.ORM.Joins("JOIN projects ON projects.id = project_users.project_id").
		Where("projects.id = ? AND project_users.user_id = ?", projectID, userID).First(&member).Error
	if err != nil {
		return nil, err
	}
	return &member, nil
}

func getTeamMember(userID, teamID uint) (*models.TeamUser, error) {
	var member models.TeamUser
	err := db.ORM.Joins("JOIN teams ON teams.id = team_users.team_id").
		Where("teams.id = ? AND team_users.user_id = ?", teamID, userID).First(&member).Error
	if err != nil {
		return nil, err
	}
	return &member, nil
}
