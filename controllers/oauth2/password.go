package oauth2controller

import (
	"strconv"

	"github.com/pinem/server/models/users"
	"gopkg.in/oauth2.v3/errors"
)

func passAuthorizationHandler(username, password string) (string, error) {
	user, err := users.Authenticate(username, password)
	if err != nil {
		return "", errors.ErrInvalidClient
	}
	return strconv.FormatUint(uint64(user.ID), 10), nil
}
