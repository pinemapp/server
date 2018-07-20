package messages

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pinem/server/utils/locale"
)

type Messages struct {
	c      *gin.Context
	t      *locale.Translate
	Errors map[string][]string
}

const messagesKey = "pinem.messages"

func GetMessages(c *gin.Context) *Messages {
	T := locale.Get(c)
	if m, ok := c.Get(messagesKey); ok {
		msg := m.(*Messages)
		msg.c = c
		msg.t = T
		return msg
	}
	msg := &Messages{c, T, make(map[string][]string)}
	c.Set(messagesKey, msg)
	return msg
}

func (msg *Messages) AddError(key, message string) {
	msg.Errors[key] = append(msg.Errors[key], message)
	msg.setInContext()
}

func (msg *Messages) AddErrorT(key, tkey string) {
	msg.AddError(key, msg.t.T(tkey))
}

func (msg *Messages) AddErrorTf(key, tkey string, args ...interface{}) {
	msg.AddError(key, fmt.Sprintf(msg.t.T(tkey), args...))
}

func (msg *Messages) ErrorT(key string, err error) {
	msg.AddErrorT(key, err.Error())
}

func (msg *Messages) GetAllErrors() map[string][]string {
	return msg.Errors
}

func (msg *Messages) GetError(key string) []string {
	if errs, ok := msg.Errors[key]; ok {
		return errs
	}
	return nil
}

func (msg *Messages) HasErrors() bool {
	return len(msg.Errors) > 0
}

func (msg *Messages) setInContext() {
	msg.c.Set(messagesKey, msg)
}
