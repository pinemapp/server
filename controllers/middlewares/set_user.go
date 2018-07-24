package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/pinem/server/controllers/router"
	"github.com/pinem/server/db"
	"github.com/pinem/server/models"
	"github.com/pinem/server/utils/auth"
)

func SetUser() gin.HandlerFunc {
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
		c.Next()
	}
}
