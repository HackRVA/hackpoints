package models

type Bounty struct {
	ID           string
	Title        string
	Description  string
	Endorsements []Member
	IsOpen       bool
}
