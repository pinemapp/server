package users

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/pinem/server/db"
	"github.com/pinem/server/errors"
	"github.com/pinem/server/models"
	"github.com/pinem/server/models/clients"
	"github.com/pinem/server/utils/messages"
	"github.com/pinem/server/utils/validators"
	"github.com/pinem/server/utils/validators/users"
	"golang.org/x/crypto/bcrypt"
)

func Create(c *gin.Context, msg *messages.Messages) (*models.User, error) {
	var f uservalidator.RegisterForm
	c.Bind(&f)

	err := validators.Validate(&f, msg)
	if err != nil {
		return nil, err
	}

	encrypted, err := encryptPassword(f.Password)
	if err != nil {
		msg.ErrorT("message", errors.ErrInternalServer)
		return nil, errors.ErrInternalServer
	}

	user := models.User{
		Name:     f.Name,
		Email:    f.Email,
		Password: encrypted,
	}

	err = db.Transaction(db.ORM, func(tx *gorm.DB) error {
		if err := tx.Create(&user).Error; err != nil {
			return err
		}
		if err := createOAuth2Client(tx, &user); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func encryptPassword(pass string) (string, error) {
	buf, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}

func createOAuth2Client(tx *gorm.DB, user *models.User) error {
	if _, err := clients.Create(tx, user); err != nil {
		return err
	}
	return nil
}
