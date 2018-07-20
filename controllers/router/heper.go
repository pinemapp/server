package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pinem/server/errors"
	"github.com/pinem/server/utils/locale"
	"github.com/pinem/server/utils/messages"
)

func RenderApiReponse(err error, c *gin.Context, renderer func()) {
	if err != nil {
		RenderError(err, c)
		return
	}
	renderer()
}

func RenderError(err error, c *gin.Context) {
	switch err {
	case errors.ErrInternalServer:
		RenderInternalServer(c)
		return
	case errors.ErrRecordNotFound:
		RenderNotFound(c)
		return
	case errors.ErrForbidden:
		RenderForbidden(c)
		return
	}

	msg := messages.GetMessages(c)
	if msg.HasErrors() {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": msg.GetAllErrors(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{
		"code":   err.Error(),
		"error":  err.Error(),
		"status": http.StatusInternalServerError,
	})
	c.Abort()
}

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
