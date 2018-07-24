package projectcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pinem/server/controllers/router"
	"github.com/pinem/server/models/projects"
	"github.com/pinem/server/utils/messages"
)

func GetProjectsHandler(c *gin.Context) {
	bs, err := projects.GetAllForUser(c)
	router.RenderApiReponse(err, c, func() {
		c.JSON(http.StatusOK, gin.H{"projects": projects.GetSimpleProjects(bs)})
	})
}

func GetProjectHandler(c *gin.Context) {
	project, err := projects.GetOneForUser(c)
	router.RenderApiReponse(err, c, func() {
		c.JSON(http.StatusOK, gin.H{"project": project})
	})
}

func PostProjectsHandler(c *gin.Context) {
	msg := messages.GetMessages(c)
	project, err := projects.Create(c, msg)
	router.RenderApiReponse(err, c, func() {
		c.JSON(http.StatusCreated, gin.H{"project": project})
	})
}

func PatchProjectHandler(c *gin.Context) {
	msg := messages.GetMessages(c)
	project, err := projects.Update(c, msg)
	router.RenderApiReponse(err, c, func() {
		c.JSON(http.StatusOK, gin.H{"project": project})
	})
}

func DeleteProjectHandler(c *gin.Context) {
	err := projects.Delete(c)
	router.RenderApiReponse(err, c, func() {
		c.Status(http.StatusOK)
	})
}
