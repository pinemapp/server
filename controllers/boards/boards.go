package boardcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pinem/server/controllers/middlewares"
	"github.com/pinem/server/controllers/router"
	"github.com/pinem/server/errors"
	"github.com/pinem/server/models/boards"
	"github.com/pinem/server/utils/messages"
)

func init() {
	r := router.Get()
	br := r.Group("/api/boards")
	br.Use(middlewares.AuthMiddleware)

	br.GET("/", GetBoardsHandler)
	br.POST("/", PostBoardsHandler)
	br.GET("/:board_id", GetBoardHandler)
	br.PATCH("/:board_id", PatchBoardHandler)
	br.DELETE("/:board_id", DeleteBoardHandler)
}

func GetBoardsHandler(c *gin.Context) {
	boards, err := boards.GetAllForUser(c)
	if err != nil {
		msg := messages.GetMessages(c)
		msg.ErrorT("message", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"errors": msg.GetAllErrors(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"boards": boards})
}

func GetBoardHandler(c *gin.Context) {
	board, err := boards.GetOneForUser(c)
	if err != nil {
		if err == errors.ErrNotFound {
			router.RenderNotFound(c)
			return
		}
		router.RenderInternalServer(c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"boards": board})
}

func PostBoardsHandler(c *gin.Context) {
	msg := messages.GetMessages(c)
	board, err := boards.Create(c, msg)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": msg.GetAllErrors(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"board": board})
}

func PatchBoardHandler(c *gin.Context) {
	msg := messages.GetMessages(c)
	board, err := boards.Update(c, msg)
	if err != nil {
		if err == errors.ErrNotFound {
			router.RenderNotFound(c)
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": msg.GetAllErrors(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"board": board})
}

func DeleteBoardHandler(c *gin.Context) {
	err := boards.Delete(c)
	if err != nil {
		if err == errors.ErrNotFound {
			router.RenderNotFound(c)
			return
		}
		router.RenderInternalServer(c)
		return
	}

	c.Status(http.StatusOK)
}
