package api

import (
	"hackpoints/api/user"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type API struct {
	UserServer *user.UserServer
}

// Setup - setup the web server
func Setup(a API) {
	r := mux.NewRouter()

	setupRoutes(a, r)

	srv := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:3000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Debug("Server listening on http://localhost:3000/")
	log.Fatal(srv.ListenAndServe())
}
