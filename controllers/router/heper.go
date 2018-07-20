package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pinem/server/errors"
	"github.com/pinem/server/utils/locale"
)

func RenderForbidden(c *gin.Context) {
	t := locale.Get(c)
	key := errors.ErrForbidden.Error()
	c.JSON(http.StatusForbidden, gin.H{
		"code":   key,
		"error":  t.T(key),
		"status": http.StatusForbidden,
	})
	c.Abort()
}

func RenderNotFound(c *gin.Context) {
	t := locale.Get(c)
	key := errors.ErrRecordNotFound.Error()
	c.JSON(http.StatusNotFound, gin.H{
		"code":   key,
		"error":  t.T(key),
		"status": http.StatusNotFound,
	})
	c.Abort()
}

func RenderInternalServer(c *gin.Context) {
	t := locale.Get(c)
	key := errors.ErrInternalServer.Error()
	c.JSON(http.StatusInternalServerError, gin.H{
		"code":   key,
		"error":  t.T(key),
		"status": http.StatusInternalServerError,
	})
	c.Abort()
}
