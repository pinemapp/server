package controllers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/pinem/server/controllers/oauth2"
	"github.com/pinem/server/controllers/router"
	_ "github.com/pinem/server/controllers/users"
)

var Router *gin.Engine

func init() {
	Router = router.Get()
}
