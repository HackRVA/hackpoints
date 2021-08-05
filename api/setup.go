package api

import (
	"hackpoints/api/auth"
	"hackpoints/api/user"
	"hackpoints/datastore/in_memory"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type API struct {
	userServer *user.UserServer
}

// Setup - setup the web server
func Setup() {
	r := mux.NewRouter()

	setupRoutes(API{
		&user.UserServer{
			Store: &in_memory.InMemoryUserStore{},
			Auth:  auth.Setup(&in_memory.InMemoryUserStore{}),
		},
	}, r)

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
