package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func setupRoutes(api API, r *mux.Router) {
	authedRoutes := r.PathPrefix("/api/").Subrouter()
	authedRoutes.Use(api.UserServer.Auth.AuthMiddleware)

	r.HandleFunc("/api/info", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		j, _ := json.Marshal(struct{ Message string }{
			Message: "hello, world!",
		})
		w.Write(j)
	})
	authedRoutes.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		j, _ := json.Marshal(struct{ Message string }{
			Message: "hello, world!",
		})
		w.Write(j)
	})

	authedRoutes.HandleFunc("/user/", api.UserServer.GetUser)
	authedRoutes.HandleFunc("/auth/login", api.UserServer.Login)
	authedRoutes.HandleFunc("/auth/register", api.UserServer.Register)

	r.HandleFunc("/login", serveLogin)

	http.Handle("/", r)
}
