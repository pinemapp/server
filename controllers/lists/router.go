package listcontroller

import (
	"github.com/pinem/server/controllers/middlewares"
	"github.com/pinem/server/controllers/router"
)

func init() {
	r := router.Get()
	middlewares.Apply(r)

	lr := r.Group("/api/projects/:project_id/lists")
	lr.Use(middlewares.Authorizer())
	{
		lr.GET("", GetListsHandler)
		lr.POST("", PostListsHandler)
		lr.GET("/:list_id", GetListHandler)
		lr.PUT("/:list_id", PatchListHandler)
		lr.PATCH("/:list_id", PatchListHandler)
		lr.DELETE("/:list_id", DeleteListHandler)
	}
}
