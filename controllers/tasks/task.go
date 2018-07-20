package taskcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pinem/server/controllers/router"
	"github.com/pinem/server/models/tasks"
	"github.com/pinem/server/utils/messages"
)

func GetTasksHandler(c *gin.Context) {
	ts, err := tasks.GetAllInBoard(c)
	router.RenderApiReponse(err, c, func() {
		c.JSON(http.StatusOK, gin.H{"tasks": ts})
	})
}

func GetTaskHandler(c *gin.Context) {
	task, err := tasks.GetOneInBoard(c)
	router.RenderApiReponse(err, c, func() {
		c.JSON(http.StatusOK, gin.H{"task": task})
	})
}

func PostTasksHandler(c *gin.Context) {
	msg := messages.GetMessages(c)
	task, err := tasks.Create(c, msg)
	router.RenderApiReponse(err, c, func() {
		c.JSON(http.StatusCreated, gin.H{"task": task})
	})
}

func PatchTaskHandler(c *gin.Context) {
	msg := messages.GetMessages(c)
	task, err := tasks.Update(c, msg)
	router.RenderApiReponse(err, c, func() {
		c.JSON(http.StatusOK, gin.H{"task": task})
	})
}

func DeleteTaskHandler(c *gin.Context) {
	err := tasks.Delete(c)
	router.RenderApiReponse(err, c, func() {
		c.Status(http.StatusNoContent)
	})
}
