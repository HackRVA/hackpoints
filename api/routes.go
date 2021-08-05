package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func setupRoutes(api API, r *mux.Router) {
	authedRoutes := r.PathPrefix("/api/").Subrouter()
	authedRoutes.Use(api.userServer.Auth.AuthMiddleware)

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

	authedRoutes.HandleFunc("/user/", api.userServer.GetUser)
	authedRoutes.HandleFunc("/auth/login", api.userServer.Login)
	authedRoutes.HandleFunc("/auth/register", api.userServer.Register)

	r.HandleFunc("/login", serveLogin)

	http.Handle("/", r)
}
