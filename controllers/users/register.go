package usercontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pinem/server/models/users"
	"github.com/pinem/server/utils/messages"
)

func PostUsersHandler(c *gin.Context) {
	msg := messages.GetMessages(c)
	user, err := users.Create(c, msg)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": msg.GetAllErrors(),
		})
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"user": user,
		})
	}
}
