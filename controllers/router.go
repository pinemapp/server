package controllers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/pinem/server/controllers/lists"
	_ "github.com/pinem/server/controllers/members"
	_ "github.com/pinem/server/controllers/projects"
	"github.com/pinem/server/controllers/router"
	_ "github.com/pinem/server/controllers/tasks"
	_ "github.com/pinem/server/controllers/teams"
	_ "github.com/pinem/server/controllers/token"
	_ "github.com/pinem/server/controllers/users"
)

var Router *gin.Engine

func init() {
	Router = router.Get()
}
