package router

import (
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/pinem/server/config"
)

var (
	once   sync.Once
	engine *gin.Engine
)

func Get() *gin.Engine {
	once.Do(func() {
		conf := config.Get()

		engine = gin.New()
		engine.Use(gin.Recovery())

		if conf.ENV == "production" {
			gin.SetMode(gin.ReleaseMode)
		} else {
			engine.Use(gin.Logger())
		}
	})
	return engine
}
