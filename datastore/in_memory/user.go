package in_memory

import (
	"errors"
	"hackpoints/models"
	"strings"
)

type InMemoryUserStore struct{}

var users = map[string]models.Credentials{
	"test": {
		Email:    "test",
		Password: "test",
	},
}

func (mem InMemoryUserStore) GetMemberByEmail(email string) (models.Member, error) {
	if val, ok := users[strings.ToLower(email)]; !ok {
		return models.Member{Email: val.Email}, errors.New("not a valid member email")
	}

	return models.Member{Email: email}, nil
}

func (mem InMemoryUserStore) SignIn(username, password string) error {
	if users[username].Password != password {
		return errors.New("error signing in")
	}

	return nil
}

func (mem InMemoryUserStore) RegisterUser(creds models.Credentials) error {
	if len(creds.Email) > 2 {
		users[creds.Email] = models.Credentials{
			Email:    creds.Email,
			Password: creds.Password,
		}
	}

	if _, ok := users[creds.Email]; ok {
		return errors.New("error registering user")
	}
	return nil
}