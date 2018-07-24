package taskcontroller

import (
	"github.com/pinem/server/controllers/middlewares"
	"github.com/pinem/server/controllers/router"
)

func init() {
	r := router.Get()
	middlewares.Apply(r)

	tr := r.Group("/api/boards/:board_id/tasks")
	tr.Use(middlewares.Authorizer())
	{
		tr.GET("", GetTasksHandler)
		tr.POST("", PostTasksHandler)
		tr.GET("/:task_id", GetTaskHandler)
		tr.PUT("/:task_id", PatchTaskHandler)
		tr.PATCH("/:task_id", PatchTaskHandler)
		tr.DELETE("/:task_id", DeleteTaskHandler)
	}
}
