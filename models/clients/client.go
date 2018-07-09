package clients

import (
	"github.com/pinem/server/db"
	"github.com/pinem/server/models"
)

func GetByUsername(username string) (*models.Client, error) {
	var client models.Client
	err := db.ORM.Joins("JOIN users ON clients.user_id = users.id").
		Where("users.name = ? OR users.email = ?", username, username).Find(&client).Error
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func GetByID(id string) (*models.Client, error) {
	var client models.Client
	err := db.ORM.Where("client_id = ?", id).First(&client).Error
	if err != nil {
		return nil, err
	}
	return &client, err
}
