package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pinem/server/errors"
	"github.com/pinem/server/utils/locale"
)

func RenderForbidden(c *gin.Context) {
	t := locale.Get(c)
	c.JSON(http.StatusForbidden, gin.H{
		"error": t.T(errors.ErrForbidden.Error()),
	})
	c.Abort()
}
