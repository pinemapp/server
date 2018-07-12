package tokencontroller

import (
	"github.com/pinem/server/controllers/middlewares"
	"github.com/pinem/server/controllers/router"
)

func init() {
	r := router.Get()
	middlewares.Apply(r)

	r.POST("/token", PostTokenHandler)
	r.GET("/verify_token", GetVerifyTokenHandler)
}
