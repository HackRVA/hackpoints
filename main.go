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
	api.Setup(api.API{
		UserServer: &user.UserServer{
			Store: &in_memory.InMemoryUserStore{},
			Auth:  auth.Setup(&in_memory.InMemoryUserStore{}),
		},
		BountyServer: &bounty.BountyServer{
			Store: &in_memory.InMemoryBountyStore{},
		},
		ScoreServer: &score.ScoreServer{
			Store: &in_memory.InMemoryScoreStore{
				BountyStore: &in_memory.InMemoryBountyStore{},
			},
		},
	})
}
