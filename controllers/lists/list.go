package listcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pinem/server/controllers/router"
	"github.com/pinem/server/errors"
	"github.com/pinem/server/models/lists"
	"github.com/pinem/server/utils/messages"
)

func GetListsHandler(c *gin.Context) {
	ls, err := lists.GetAllInBoard(c)
	if err != nil {
		router.RenderNotFound(c)
		return
	}
	c.JSON(http.StatusOK, gin.H{"lists": lists.GetSimpleLists(ls)})
}

func GetListHandler(c *gin.Context) {
	list, err := lists.GetOneInBoard(c)
	if err != nil {
		router.RenderNotFound(c)
		return
	}
	c.JSON(http.StatusOK, gin.H{"list": list})
}

func PostListsHandler(c *gin.Context) {
	msg := messages.GetMessages(c)
	list, err := lists.Create(c, msg)
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

	c.JSON(http.StatusCreated, gin.H{"list": list})
}

func PatchListHandler(c *gin.Context) {
	msg := messages.GetMessages(c)
	list, err := lists.Update(c, msg)
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
	c.JSON(http.StatusOK, gin.H{"list": list})
}

func DeleteListHandler(c *gin.Context) {
	err := lists.Delete(c)
	if err != nil {
		router.RenderNotFound(c)
		return
	}
	c.Status(http.StatusNoContent)
}
