package boardcontroller

import (
	"github.com/pinem/server/controllers/middlewares"
	"github.com/pinem/server/controllers/router"
)

func init() {
	r := router.Get()
	middlewares.Apply(r)

	br := r.Group("/api/boards")
	{
		br.GET("", GetBoardsHandler)
		br.POST("", PostBoardsHandler)
		br.GET("/:board_id", GetBoardHandler)
		br.PUT("/:board_id", PatchBoardHandler)
		br.PATCH("/:board_id", PatchBoardHandler)
		br.DELETE("/:board_id", DeleteBoardHandler)
	}
}
