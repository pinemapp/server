package tokencontroller

import (
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/pinem/server/config"
	"github.com/pinem/server/controllers/router"
	"github.com/pinem/server/models"
	"github.com/pinem/server/models/users"
	"github.com/pinem/server/utils/auth"
	"github.com/pinem/server/utils/locale"
)

type credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func PostTokenHandler(c *gin.Context) {
	var cred credentials
	c.Bind(&cred)

	t := locale.Get(c)
	user, err := users.Authenticate(cred.Username, cred.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": t.T(err.Error()),
		})
		return
	}

	token, err := generateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": t.T("errors_internal_server"),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func GetVerifyTokenHandler(c *gin.Context) {
	_, err := auth.Verify(c)
	if err != nil {
		router.RenderForbidden(c)
		return
	}
	c.Status(http.StatusOK)
}

func generateToken(user *models.User) (string, error) {
	conf := config.Get()
	now := time.Now().Unix()
	claims := models.Claims{
		models.UserClaims{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
		jwt.StandardClaims{
			Issuer:    conf.App,
			ExpiresAt: now + conf.Token.ExpiresAt,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(conf.Secret))
}
