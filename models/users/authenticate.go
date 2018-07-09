package users

import (
	"github.com/pinem/server/db"
	"github.com/pinem/server/models"
	"golang.org/x/crypto/bcrypt"
)

func Authenticate(username, password string) (*models.User, error) {
	var user models.User
	if err := db.ORM.Where("name = ? OR email = ?", username, username).First(&user).Error; err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}
	return &user, nil
}
