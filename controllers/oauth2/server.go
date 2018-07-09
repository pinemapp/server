package oauth2controller

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pinem/server/config"
	"github.com/pinem/server/controllers/router"
	redis "gopkg.in/go-oauth2/redis.v1"
	"gopkg.in/oauth2.v3/errors"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/server"
)

func init() {
	conf := config.Get()

	m := manage.NewDefaultManager()
	m.MustTokenStorage(redis.NewTokenStore(&redis.Config{
		Addr:     fmt.Sprintf("%s:%d", conf.Redis.Host, conf.Redis.Port),
		Password: conf.Redis.Password,
		DB:       conf.Redis.Db,
	}))

	clientStore := &ClientStore{}
	m.MapClientStorage(clientStore)

	tokenExp := time.Duration(conf.OAuth2.AccessTokenExp) * time.Second
	refreshExp := time.Duration(conf.OAuth2.RefreshTokenExp) * time.Second
	tokenConf := &manage.Config{
		AccessTokenExp:    tokenExp,
		RefreshTokenExp:   refreshExp,
		IsGenerateRefresh: true,
	}
	m.SetPasswordTokenCfg(tokenConf)
	m.SetAuthorizeCodeTokenCfg(tokenConf)
	m.SetRefreshTokenCfg(&manage.RefreshingConfig{
		AccessTokenExp:     tokenExp,
		RefreshTokenExp:    refreshExp,
		IsGenerateRefresh:  true,
		IsRemoveAccess:     true,
		IsRemoveRefreshing: true,
	})

	s := server.NewDefaultServer(m)
	s.SetClientInfoHandler(clientInfoHandler)
	s.SetInternalErrorHandler(internalErrorHandler)
	s.SetResponseErrorHandler(responseErrorHandler)
	s.SetPasswordAuthorizationHandler(passAuthorizationHandler)

	r := router.Get()
	r.POST("/token", gin.WrapF(func(res http.ResponseWriter, req *http.Request) {
		s.HandleTokenRequest(res, req)
	}))
}

func internalErrorHandler(err error) (r *errors.Response) {
	log.Println(err.Error())
	return
}

func responseErrorHandler(r *errors.Response) {
	log.Println(r.Error.Error())
	return
}
