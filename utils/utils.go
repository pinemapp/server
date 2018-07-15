package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetIntParam(key string, c *gin.Context) uint {
	s := c.Param(key)
	id, err := strconv.ParseUint(s, 10, 0)
	if err != nil {
		return 0
	}
	return uint(id)
}

func GetOrderRange(newOrder, oldOrder int) (min, max int) {
	if oldOrder < newOrder {
		min = oldOrder - 1
		max = newOrder
	} else {
		min = newOrder - 1
		max = oldOrder + 1
	}
	return
}
