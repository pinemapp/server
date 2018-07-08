package usercontroller

import "github.com/pinem/server/controllers/router"

func init() {
	r := router.Get()
	ur := r.Group("/api/users")
	ur.POST("/", PostUsersHandler)
}
