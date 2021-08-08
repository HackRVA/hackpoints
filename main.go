package main

import (
	"hackpoints/api"
	"hackpoints/api/auth"
	"hackpoints/api/bounty"
	"hackpoints/api/score"
	"hackpoints/api/user"
	"hackpoints/datastore/in_memory"
)

func main() {
	datastore := &in_memory.Store{}
	api.Setup(api.API{
		UserServer: &user.UserServer{
			Store: datastore,
			Auth:  auth.Setup(datastore),
		},
		BountyServer: &bounty.BountyServer{},
		ScoreServer: &score.ScoreServer{
			Store: datastore,
		},
	})
}
