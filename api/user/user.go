package user

import (
	"encoding/json"
	"errors"
	"hackpoints/api/auth"
	"hackpoints/datastore"
	"hackpoints/models"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
)

type UserServer struct {
	Store datastore.UserStore
	Auth  auth.AuthProvider
}

func (u *UserServer) GetUser(w http.ResponseWriter, r *http.Request) {
	userProfile, err := u.Store.GetMemberByEmail(u.Auth.User(r).GetUserName())
	if err != nil {
		http.Error(w, "user not found", http.StatusUnauthorized)
		return
	}

	if userProfile.Email == (models.UserResponse{}).Email {
		http.Error(w, "user not found", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	j, _ := json.Marshal(userProfile)
	w.Write(j)
}

func (u *UserServer) Register(w http.ResponseWriter, r *http.Request) {
	// Parse and decode the request body into a new `Credentials` instance
	creds := &models.Credentials{}
	err := json.NewDecoder(r.Body).Decode(creds)
	if err != nil {
		// If there is something wrong with the request body, return a 400 status
		http.Error(w, "error registering user", http.StatusBadRequest)
		return
	}

	if len(creds.Password) < 3 {
		http.Error(w, "password must be longer", http.StatusBadRequest)
		return
	}

	err = u.Store.RegisterUser(models.Credentials{
		Email:    strings.ToLower(creds.Email),
		Password: creds.Password,
	})

	if err != nil {
		log.Error(err)
		http.Error(w, "error registering user", http.StatusBadRequest)
		return
	}

	// We reach this point if the credentials we correctly stored in the database, and the default status of 200 is sent back
	w.Header().Set("Content-Type", "application/json")
	j, _ := json.Marshal(models.EndpointSuccess{
		Ack: true,
	})
	w.Write(j)
}

func (u *UserServer) Login(w http.ResponseWriter, r *http.Request) {
	log.Debug("attempting to login")
	user := u.Auth.User(r)
	if user == nil {
		http.Error(w, errors.New("user is nil???").Error(), http.StatusUnprocessableEntity)
		return
	}

	token, err := u.Auth.IssueAccessToken(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	tokenJSON := &models.TokenResponse{}
	tokenJSON.Token = token

	w.Header().Set("Content-Type", "application/json")

	j, _ := json.Marshal(tokenJSON)
	w.Write(j)
}
