package teamcontroller

import (
	"github.com/pinem/server/controllers/middlewares"
	"github.com/pinem/server/controllers/router"
)

func init() {
	r := router.Get()
	middlewares.Apply(r)

	tr := r.Group("/api/teams")
	{
		tr.GET("", GetTeamsHandler)
		tr.POST("", PostTeamsHandler)
		tr.GET("/:team_id", GetTeamHandler)
		tr.PUT("/:team_id", PatchTeamHandler)
		tr.PATCH("/:team_id", PatchTeamHandler)
		tr.DELETE("/:team_id", DeleteTeamHandler)
	}
}
