//     Version: 0.0.1
//     License: MIT http://opensource.org/licenses/MIT
//
//    Schemes: https, http
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
	"hackpoints/models"
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
	//   Version and commit hash
	//
	//     Produces:
	//     - application/json
	//
	//     Responses:
	//       200: infoResponse
	r.HandleFunc("/api/info", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		j, _ := json.Marshal(models.InfoResponse{
			Message: "hello, world!",
			Version: 0,
			Commit:  "demo",
		})
		w.Write(j)
	}).Methods(http.MethodGet)
	// swagger:route GET /api/score score score
	//
	// Shows the current score for the space
	//
	//   Score is tallied by adding up the number of endorsements on closed bounties
	//
	//     Produces:
	//     - application/json
	//
	//     Security:
	//     - basicAuth:
	//
	//     Responses:
	//       200: scoreResponse
	authedRoutes.HandleFunc("/score", api.ScoreServer.Get).Methods(http.MethodGet)
	// swagger:route GET /api/bounty bounty bounty
	//
	// Retrieves many bounties.
	//
	//     Produces:
	//     - application/json
	//
	//     Security:
	//     - basicAuth:
	//
	//     Responses:
	//       200: bountyResponse

	// swagger:route POST /api/bounty bounty bountyGetRequest
	//
	// Retrieves one bounty.
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
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
	//   A bounty is a task or action item that one person or a group of people can complete.
	//   Members will decide to endorse certain bounties.
	//   When the bounty is closed, the number of endorsements on that bounty gets added to the
	//   groups total score.  At certain score intervals, we have a pizza party.
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Security:
	//     - basicAuth:
	//
	//     Responses:
	//       200: endpointSuccessResponse
	authedRoutes.HandleFunc("/bounty/new", api.BountyServer.New).Methods(http.MethodPost)
	// swagger:route PATCH /api/bounty/endorse bounty bountyEndorseRequest
	//
	// Endorse a bounty
	//
	//  An endorsement is basically assigning one point value to the bounty.
	//  A member can only endorse a bounty once.
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Security:
	//     - basicAuth:
	//
	//     Responses:
	//       200: endpointSuccessResponse
	authedRoutes.HandleFunc("/bounty/endorse", api.BountyServer.Endorse).Methods(http.MethodPatch)
	// swagger:route PATCH /api/bounty/close bounty bountyCloseRequest
	//
	// Close a bounty
	//
	//   When a bounty is completed, it's no longer available for people to endorse.
	//   All closed bounties will have the number of endorsements added to the Space's total score.
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//
	//     Security:
	//     - basicAuth:
	//
	//     Responses:
	//       200: endpointSuccessResponse
	authedRoutes.HandleFunc("/bounty/close", api.BountyServer.Endorse).Methods(http.MethodPatch)
	// swagger:route GET /api/user user user
	//
	// Shows the current logged in user
	//
	//   retrieve user information so that we can easily display it in UIs
	//
	//     Produces:
	//     - application/json
	//
	//     Security:
	//     - basicAuth:
	//
	//     Responses:
	//       200: userResponse
	authedRoutes.HandleFunc("/user", api.UserServer.GetUser).Methods(http.MethodGet)
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
	//  Eventually this will verify that only a valid member can sign up.
	//  Currently this endpoint doesn't work until we get a database.
	//
	//     Produces:
	//     - application/json
	//
	//
	//     Responses:
	//       200: endpointSuccessResponse
	r.HandleFunc("/api/auth/register", api.UserServer.Register).Methods(http.MethodPost)

	http.Handle("/", r)
}
