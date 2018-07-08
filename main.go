package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/pinem/server/config"
	"github.com/pinem/server/controllers"
	"github.com/pinem/server/db"
)

func main() {
	conf := config.Get()

	if err := db.InitORM(); err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}

	server := &http.Server{
		Addr:         conf.ServerAddr(),
		Handler:      controllers.Router,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	log.Printf("Loading %s environment\n", conf.ENV)
	log.Printf("Listening to %s\n", conf.ServerAddr())
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}
}
