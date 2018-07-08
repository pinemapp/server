package tokencontroller

import (
	oauth2 "gopkg.in/oauth2.v3"
)

type ClientStore struct {
}

func newClientStore() *ClientStore {
	return &ClientStore{}
}

func (s *ClientStore) GetByID(id string) (oauth2.ClientInfo, error) {
	return nil, nil
}
