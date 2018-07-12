package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pinem/server/errors"
	"github.com/pinem/server/utils/locale"
)

func NotFound() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Writer.Status() == 404 {
			t := locale.Get(c)
			key := errors.ErrNotFound.Error()
			c.JSON(http.StatusNotFound, gin.H{
				"code":   key,
				"error":  t.T(key),
				"status": http.StatusNotFound,
			})
			c.Abort()
		}
	}
}
