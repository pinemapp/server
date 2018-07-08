package tokencontroller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pinem/server/config"
	"github.com/pinem/server/controllers/router"
	"github.com/pinem/server/models/users"
	redis "gopkg.in/go-oauth2/redis.v1"
	"gopkg.in/oauth2.v3/errors"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/server"
)

type credentials struct {
	Username string `json:"username"`
	Password string `josn:"password"`
}

func init() {
	conf := config.Get()

	m := manage.NewDefaultManager()
	m.MustTokenStorage(redis.NewTokenStore(&redis.Config{
		Addr:     fmt.Sprintf("%s:%d", conf.Redis.Host, conf.Redis.Port),
		Password: conf.Redis.Password,
		DB:       conf.Redis.Db,
	}))

	clientStore := newClientStore()
	m.MapClientStorage(clientStore)

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

func clientInfoHandler(req *http.Request) (string, string, error) {
	// var cred credentials
	// dec := json.NewDecoder(req.Body)

	// if err := dec.Decode(&cred); err != nil {
	// 	return "", "", err
	// }
	// return cred.Username, cred.Password, nil
	username, password := req.FormValue("username"), req.FormValue("password")
	return username, password, nil
}

func internalErrorHandler(err error) (r *errors.Response) {
	log.Println(err.Error())
	return
}

func responseErrorHandler(r *errors.Response) {
	log.Println(r.Error.Error())
	return
}

func passAuthorizationHandler(username, password string) (string, error) {
	user, err := users.Authenticate(username, password)
	if err != nil {
		return "", err
	}
	return strconv.FormatUint(uint64(user.ID), 10), nil
}
