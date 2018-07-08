package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/pinem/server/controllers/router"
	_ "github.com/pinem/server/controllers/users"
)

var Router *gin.Engine

func init() {
	Router = router.Get()
}
