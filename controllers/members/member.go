package membercontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pinem/server/controllers/router"
	"github.com/pinem/server/errors"
	"github.com/pinem/server/models/members"
	"github.com/pinem/server/utils/messages"
)

func PostMembersHandler(c *gin.Context) {
	msg := messages.GetMessages(c)
	member, err := members.Add(c, msg)
	if err != nil {
		if err == errors.ErrRecordNotFound {
			router.RenderNotFound(c)
			return
		}
		if err == errors.ErrMemberAlreadyInBoard || err == errors.ErrUserNotExist {
			msg.ErrorT("user_id", err)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": msg.GetAllErrors(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"member": member})
}

func PatchMemberHandler(c *gin.Context) {
	msg := messages.GetMessages(c)
	member, err := members.Update(c, msg)
	if err != nil {
		if err == errors.ErrRecordNotFound {
			router.RenderNotFound(c)
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": msg.GetAllErrors(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"member": member})
}

func DeleteMemberHandler(c *gin.Context) {
	if err := members.Delete(c); err != nil {
		if err == errors.ErrRecordNotFound {
			router.RenderNotFound(c)
			return
		}
		router.RenderInternalServer(c)
		return
	}

	c.Status(http.StatusOK)
}
