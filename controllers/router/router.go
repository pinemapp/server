package router

import (
	"sync"

	"github.com/casbin/casbin"
	"github.com/gin-gonic/gin"
	"github.com/pinem/server/config"
)

var (
	once   sync.Once
	engine *gin.Engine

	once1    sync.Once
	enforcer *casbin.Enforcer
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

func GetEnforcer() *casbin.Enforcer {
	once1.Do(func() {
		enforcer = casbin.NewEnforcer("./config/auth_model.conf", "./config/policy.csv")
	})
	return enforcer
}
