package taskcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pinem/server/controllers/router"
	"github.com/pinem/server/errors"
	"github.com/pinem/server/models/tasks"
	"github.com/pinem/server/utils/messages"
)

func GetTasksHandler(c *gin.Context) {
	ts, err := tasks.GetAllInBoard(c)
	if err != nil {
		router.RenderNotFound(c)
		return
	}
	c.JSON(http.StatusOK, gin.H{"tasks": ts})
}

func GetTaskHandler(c *gin.Context) {
	task, err := tasks.GetOneInBoard(c)
	if err != nil {
		router.RenderNotFound(c)
		return
	}
	c.JSON(http.StatusOK, gin.H{"task": task})
}

func PostTasksHandler(c *gin.Context) {
	msg := messages.GetMessages(c)
	task, err := tasks.Create(c, msg)
	if err != nil {
		if err == errors.ErrNotFound {
			router.RenderNotFound(c)
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": msg.GetAllErrors(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"task": task})
}

func PatchTaskHandler(c *gin.Context) {
	msg := messages.GetMessages(c)
	task, err := tasks.Update(c, msg)
	if err != nil {
		if err == errors.ErrNotFound {
			router.RenderNotFound(c)
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": msg.GetAllErrors(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"task": task})
}

func DeleteTaskHandler(c *gin.Context) {
	err := tasks.Delete(c)
	if err != nil {
		router.RenderNotFound(c)
		return
	}
	c.Status(http.StatusNoContent)
}
