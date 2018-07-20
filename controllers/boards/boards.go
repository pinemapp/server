package boardcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pinem/server/controllers/router"
	"github.com/pinem/server/models/boards"
	"github.com/pinem/server/utils/messages"
)

func GetBoardsHandler(c *gin.Context) {
	bs, err := boards.GetAllForUser(c)
	router.RenderApiReponse(err, c, func() {
		c.JSON(http.StatusOK, gin.H{"boards": boards.GetSimpleBoards(bs)})
	})
}

func GetBoardHandler(c *gin.Context) {
	board, err := boards.GetOneForUser(c)
	router.RenderApiReponse(err, c, func() {
		c.JSON(http.StatusOK, gin.H{"boards": board})
	})
}

func PostBoardsHandler(c *gin.Context) {
	msg := messages.GetMessages(c)
	board, err := boards.Create(c, msg)
	router.RenderApiReponse(err, c, func() {
		c.JSON(http.StatusCreated, gin.H{"board": board})
	})
}

func PatchBoardHandler(c *gin.Context) {
	msg := messages.GetMessages(c)
	board, err := boards.Update(c, msg)
	router.RenderApiReponse(err, c, func() {
		c.JSON(http.StatusOK, gin.H{"board": board})
	})
}

func DeleteBoardHandler(c *gin.Context) {
	err := boards.Delete(c)
	router.RenderApiReponse(err, c, func() {
		c.Status(http.StatusOK)
	})
}
