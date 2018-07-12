package controllers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/pinem/server/controllers/boards"
	_ "github.com/pinem/server/controllers/members"
	"github.com/pinem/server/controllers/router"
	_ "github.com/pinem/server/controllers/token"
	_ "github.com/pinem/server/controllers/users"
)

var Router *gin.Engine

func init() {
	Router = router.Get()
}
