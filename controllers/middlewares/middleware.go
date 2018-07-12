package middlewares

import (
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/pinem/server/controllers/router"
)

var once sync.Once

func Apply(engine *gin.Engine) {
	once.Do(func() {
		router.Get().Use(NotFound())
		router.Get().Use(Authorizer(router.GetEnforcer()))
	})
}
