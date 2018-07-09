package oauth2controller

import (
	"net/http"

	"github.com/pinem/server/models"
	"github.com/pinem/server/models/clients"
	oauth2 "gopkg.in/oauth2.v3"
	"gopkg.in/oauth2.v3/errors"
)

type ClientStore struct {
}

func (s *ClientStore) GetByID(id string) (oauth2.ClientInfo, error) {
	info, err := clients.GetByID(id)
	if err != nil {
		return nil, errors.ErrInvalidClient
	}
	return info, nil
}

func clientInfoHandler(req *http.Request) (string, string, error) {
	var client *models.Client
	grantType := req.FormValue("grant_type")

	switch oauth2.GrantType(grantType) {
	case oauth2.ClientCredentials:
		id := req.FormValue("client_id")
		secret := req.FormValue("client_secret")
		c, err := clients.GetByID(id)
		if err != nil {
			return "", "", errors.ErrInvalidClient
		}
		if c.ClientSecret != secret {
			return "", "", errors.ErrInvalidClient
		}
		client = c
	case oauth2.PasswordCredentials:
		username := req.FormValue("username")
		c, err := clients.GetByUsername(username)
		if err != nil {
			return "", "", errors.ErrInvalidClient
		}
		client = c
	default:
		return "", "", errors.ErrInvalidGrant
	}
	return client.ClientID, client.ClientSecret, nil
}
