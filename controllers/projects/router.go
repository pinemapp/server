package projectcontroller

import (
	"github.com/pinem/server/controllers/middlewares"
	"github.com/pinem/server/controllers/router"
)

func init() {
	r := router.Get()
	middlewares.Apply(r)

	br := r.Group("/api/projects")
	br.Use(middlewares.Authorizer())
	{
		br.GET("", GetProjectsHandler)
		br.POST("", PostProjectsHandler)
		br.GET("/:project_id", GetProjectHandler)
		br.PUT("/:project_id", PatchProjectHandler)
		br.PATCH("/:project_id", PatchProjectHandler)
		br.DELETE("/:project_id", DeleteProjectHandler)
	}
}
