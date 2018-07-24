package teamcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pinem/server/controllers/router"
	"github.com/pinem/server/models/teams"
)

func GetTeamsHandler(c *gin.Context) {
	teams, err := teams.FindAllFromContext(c)
	router.RenderApiReponse(err, c, func() {
		c.JSON(http.StatusOK, gin.H{"teams": teams})
	})
}

func GetTeamHandler(c *gin.Context) {
	team, err := teams.FindFromContext(c)
	router.RenderApiReponse(err, c, func() {
		c.JSON(http.StatusOK, gin.H{"team": team})
	})
}

func PostTeamsHandler(c *gin.Context) {
	team, err := teams.CreateFromContext(c)
	router.RenderApiReponse(err, c, func() {
		c.JSON(http.StatusCreated, gin.H{"team": team})
	})
}

func PatchTeamHandler(c *gin.Context) {
	team, err := teams.UpdateFromContext(c)
	router.RenderApiReponse(err, c, func() {
		c.JSON(http.StatusOK, gin.H{"team": team})
	})
}

func DeleteTeamHandler(c *gin.Context) {
	err := teams.DeleteFromContext(c)
	router.RenderApiReponse(err, c, func() {
		c.Status(http.StatusNoContent)
	})
}
