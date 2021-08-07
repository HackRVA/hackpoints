//     Schemes: http, https
//     Version: 0.0.1
//     License: MIT http://opensource.org/licenses/MIT
//
//    SecurityDefinitions:
//    bearerAuth:
//      type: apiKey
//      in: header
//      name: Authorization
//      description: Enter your bearer token
//    basicAuth:
//      type: basic
//      in: header
//      name: Authorization
//      description: Enter your basic auth credentials
//
// swagger:meta
package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func setupRoutes(api API, r *mux.Router) {
	authedRoutes := r.PathPrefix("/api/").Subrouter()
	authedRoutes.Use(api.UserServer.Auth.AuthMiddleware)
	// swagger:route GET /api/info info info
	//
	// Returns some debug info
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https
	//
	//     Responses:
	//       200: infoResponse
	r.HandleFunc("/info", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		j, _ := json.Marshal(struct{ Message string }{
			Message: "hello, world!",
		})
		w.Write(j)
	})
	// swagger:route GET /api/score score score
	//
	// Shows the current score for the space
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https
	//
	//     Security:
	//     - basicAuth:
	//
	//     Responses:
	//       200: scoreResponse
	authedRoutes.HandleFunc("/score", api.ScoreServer.Get).Methods(http.MethodGet)
	// swagger:route POST /api/bounty bounty bountyGetRequest
	//
	// Retrieves one or many bounties.
	//   If no bounty is submitted, it will return all bounties.
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https
	//
	//     Security:
	//     - basicAuth:
	//
	//     Responses:
	//       200: bountyResponse
	authedRoutes.HandleFunc("/bounty", api.BountyServer.Get).Methods(http.MethodPost, http.MethodGet)
	// swagger:route POST /api/bounty/new bounty bountyNewRequest
	//
	// Creates a new bounty
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https
	//
	//     Security:
	//     - basicAuth:
	//
	//     Responses:
	//       200: endpointSuccessResponse
	authedRoutes.HandleFunc("/bounty/new", api.BountyServer.New).Methods(http.MethodPost)
	// swagger:route POST /api/bounty/endorse bounty bountyEndorseRequest
	//
	// Endorse a bounty
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https
	//
	//     Security:
	//     - basicAuth:
	//
	//     Responses:
	//       200: endpointSuccessResponse
	authedRoutes.HandleFunc("/bounty/endorse", api.BountyServer.Endorse).Methods(http.MethodPost)
	// swagger:route GET /api/user user user
	//
	// Shows the current logged in user
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https
	//
	//     Security:
	//     - basicAuth:
	//
	//     Responses:
	//       200: userResponse
	authedRoutes.HandleFunc("/user", api.UserServer.GetUser)
	// swagger:route POST /api/auth/login auth auth
	//
	// Login
	//
	// Login accepts some json with the `email` and `password`
	//   and returns some json that has the token string
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https
	//
	//     security:
	//       - basicAuth: []
	//
	//     Responses:
	//       200: loginResponse
	authedRoutes.HandleFunc("/auth/login", api.UserServer.Login).Methods(http.MethodPost)
	// swagger:route POST /api/auth/register auth registerUserRequest
	//
	// Register a new user
	//
	// Register a new user of the app
	//
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https
	//
	//     Responses:
	//       200: endpointSuccessResponse
	r.HandleFunc("/api/auth/register", api.UserServer.Register).Methods(http.MethodPost)

	// r.HandleFunc("/login", serveLogin)

	http.Handle("/", r)
}
