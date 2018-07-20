package membercontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pinem/server/controllers/router"
	"github.com/pinem/server/models/members"
	"github.com/pinem/server/utils/messages"
)

func PostMembersHandler(c *gin.Context) {
	msg := messages.GetMessages(c)
	member, err := members.Add(c, msg)
	router.RenderApiReponse(err, c, func() {
		c.JSON(http.StatusCreated, gin.H{"member": member})
	})
}

func PatchMemberHandler(c *gin.Context) {
	msg := messages.GetMessages(c)
	member, err := members.Update(c, msg)
	router.RenderApiReponse(err, c, func() {
		c.JSON(http.StatusOK, gin.H{"member": member})
	})
}

func DeleteMemberHandler(c *gin.Context) {
	err := members.Delete(c)
	router.RenderApiReponse(err, c, func() {
		c.Status(http.StatusOK)
	})
}
