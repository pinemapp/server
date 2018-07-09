package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetParamID(c *gin.Context) uint {
	s := c.Param("id")
	id, err := strconv.ParseUint(s, 10, 0)
	if err != nil {
		return 0
	}
	return uint(id)
}
