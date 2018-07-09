package clients

import (
	"github.com/jinzhu/gorm"
	"github.com/pinem/server/models"
	uuid "github.com/satori/go.uuid"
)

func Create(tx *gorm.DB, user *models.User) (*models.Client, error) {
	u := uuid.Must(uuid.NewV4())
	u1 := uuid.Must(uuid.NewV4())

	client := models.Client{
		UserID:       user.ID,
		ClientID:     u.String(),
		ClientSecret: u1.String(),
	}
	if err := tx.Create(&client).Error; err != nil {
		return nil, err
	}
	return &client, nil
}
