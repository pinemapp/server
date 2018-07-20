package utils

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mgutz/str"
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

func GenerateSlug(name string) string {
	return strings.Trim(str.Dasherize(name), "-")
}
