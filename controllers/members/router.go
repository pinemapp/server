package membercontroller

import (
	"github.com/pinem/server/controllers/middlewares"
	"github.com/pinem/server/controllers/router"
)

func init() {
	r := router.Get()
	middlewares.Apply(r)

	mr := r.Group("/api/projects/:project_id/members")
	mr.Use(middlewares.Authorizer())
	{
		mr.POST("", PostMembersHandler)
		mr.PUT("/:member_id", PatchMemberHandler)
		mr.PATCH("/:member_id", PatchMemberHandler)
		mr.DELETE("/:member_id", DeleteMemberHandler)
	}
}
