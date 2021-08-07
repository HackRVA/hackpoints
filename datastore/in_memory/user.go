package in_memory

import (
	"errors"
	"hackpoints/models"
	"strings"
)

var users = map[string]models.Credentials{
	"test": {
		Email:    "test",
		Password: "test",
	},
	"test1": {
		Email:    "test1",
		Password: "test",
	},
	"test2": {
		Email:    "test2",
		Password: "test2",
	},
}

func (mem Store) GetMemberByEmail(email string) (models.Member, error) {
	if val, ok := users[strings.ToLower(email)]; !ok {
		return models.Member{Email: val.Email}, errors.New("not a valid member email")
	}

	return models.Member{Email: email}, nil
}

func (mem Store) SignIn(username, password string) error {
	if users[username].Password != password {
		return errors.New("error signing in")
	}

	return nil
}

func (mem Store) RegisterUser(creds models.Credentials) error {
	if len(creds.Email) < 3 {
		return errors.New("email too short")
	}

	if _, ok := users[creds.Email]; !ok {
		return errors.New("error registering user - not a member")
	}

	users[creds.Email] = models.Credentials{
		Email:    creds.Email,
		Password: creds.Password,
	}

	return nil
}
