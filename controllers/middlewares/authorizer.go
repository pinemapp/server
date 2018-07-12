package middlewares

import (
	"github.com/casbin/casbin"
	"github.com/gin-gonic/gin"
	"github.com/pinem/server/controllers/router"
	"github.com/pinem/server/db"
	"github.com/pinem/server/models"
	"github.com/pinem/server/utils"
	"github.com/pinem/server/utils/auth"
)

func Authorizer(e *casbin.Enforcer) func(*gin.Context) {
	return func(c *gin.Context) {
		claims, err := auth.Verify(c)
		if claims != nil && err == nil {
			var user models.User
			if err := db.ORM.First(&user, claims.User.ID).Error; err != nil {
				router.RenderForbidden(c)
				return
			}
			auth.SetUserInContext(&user, c)
		}

		user := auth.GetUserFromContext(c)
		if user == nil {
			res, err := e.EnforceSafe("anonymous", c.Request.URL.Path, c.Request.Method)
			if err != nil {
				router.RenderForbidden(c)
				return
			}
			if !res {
				router.RenderForbidden(c)
				return
			}
			return
		}

		boardID := utils.GetIntParam("board_id", c)
		if boardID == 0 {
			res, err := e.EnforceSafe("user", c.Request.URL.Path, c.Request.Method)
			if err != nil {
				router.RenderForbidden(c)
				return
			}
			if !res {
				router.RenderForbidden(c)
				return
			}
			return
		}

		var member models.BoardUser
		err = db.ORM.Joins("JOIN boards ON boards.id = board_users.board_id").
			Where("boards.id = ? AND board_users.user_id = ?", boardID, user.ID).First(&member).Error
		if err != nil {
			router.RenderForbidden(c)
			return
		}

		role := string(member.Role)
		if role == "" {
			role = "user"
		}

		res, err := e.EnforceSafe(role, c.Request.URL.Path, c.Request.Method)
		if err != nil {
			router.RenderForbidden(c)
			return
		}
		if !res {
			router.RenderForbidden(c)
			return
		}

		c.Next()
	}
}
