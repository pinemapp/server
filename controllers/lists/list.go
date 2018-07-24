package listcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pinem/server/controllers/router"
	"github.com/pinem/server/models/lists"
	"github.com/pinem/server/utils/messages"
)

func GetListsHandler(c *gin.Context) {
	ls, err := lists.GetAllInProject(c)
	router.RenderApiReponse(err, c, func() {
		c.JSON(http.StatusOK, gin.H{"lists": lists.GetSimpleLists(ls)})
	})
}

func GetListHandler(c *gin.Context) {
	list, err := lists.GetOneInProject(c)
	router.RenderApiReponse(err, c, func() {
		c.JSON(http.StatusOK, gin.H{"list": list})
	})
}

func PostListsHandler(c *gin.Context) {
	msg := messages.GetMessages(c)
	list, err := lists.Create(c, msg)
	router.RenderApiReponse(err, c, func() {
		c.JSON(http.StatusCreated, gin.H{"list": list})
	})
}

func PatchListHandler(c *gin.Context) {
	msg := messages.GetMessages(c)
	list, err := lists.Update(c, msg)
	router.RenderApiReponse(err, c, func() {
		c.JSON(http.StatusOK, gin.H{"list": list})
	})
}

func DeleteListHandler(c *gin.Context) {
	err := lists.Delete(c)
	router.RenderApiReponse(err, c, func() {
		c.Status(http.StatusNoContent)
	})
}
