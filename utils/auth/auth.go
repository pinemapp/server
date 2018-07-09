package auth

import (
	"regexp"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/pinem/server/config"
	"github.com/pinem/server/errors"
	"github.com/pinem/server/models"
)

var bearerRegex = regexp.MustCompile(`^Bearer (.*)$`)

func Verify(c *gin.Context) (*models.Claims, error) {
	content := c.GetHeader("Authorization")

	if content == "" {
		return nil, errors.ErrForbidden
	}

	m := bearerRegex.FindStringSubmatch(content)
	if len(m) != 2 {
		return nil, errors.ErrForbidden
	}

	token, err := jwt.ParseWithClaims(m[1], &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.ErrForbidden
		}
		return []byte(config.Get().Secret), nil
	})
	if err != nil {
		return nil, errors.ErrForbidden
	}

	claims, ok := token.Claims.(*models.Claims)
	if !ok || !token.Valid {
		return nil, errors.ErrForbidden
	}

	return claims, nil
}
