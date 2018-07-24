package utils

import (
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mgutz/str"
	"github.com/pinem/server/config"
)

var (
	num      = []rune("0123456789")
	alphaNum = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
)

func init() {
	rand.Seed(time.Now().In(config.Get().GetLocation()).UnixNano())
}

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

func RandomString(n int) string {
	return randomNFrom(n, alphaNum)
}

func RandomNumString(n int) string {
	return randomNFrom(n, num)
}

func randomNFrom(n int, source []rune) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = source[rand.Intn(len(source))]
	}
	return string(b)
}
