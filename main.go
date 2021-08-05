package main

import (
	"hackpoints/api"
	"hackpoints/api/auth"
	"hackpoints/api/user"
	"hackpoints/datastore/in_memory"
)

func main() {
	api.Setup(api.API{
		UserServer: &user.UserServer{
			Store: &in_memory.InMemoryUserStore{},
			Auth:  auth.Setup(&in_memory.InMemoryUserStore{}),
		},
	})
}
