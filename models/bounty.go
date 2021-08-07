package models

type BountyStore interface {
	New(Bounty) error
	Update(Bounty) error
	Get(Bounty) ([]Bounty, error)
	Endorse(Bounty, Member) error
}

// swagger:parameters bountyNewRequest
type bountyNewRequest struct {
	// in: body
	Body Bounty
}

// swagger:parameters bountyEndorseRequest
type bountyEndorseRequest struct {
	// in: body
	Body Bounty
}

// swagger:parameters bountyGetRequest
type bountyGetRequest struct {
	// in: body
	Body Bounty
}

// swagger:response bountyResponse
type bountyResponse struct {
	Body Bounty
}

type Bounty struct {
	ID           string
	Title        string
	Description  string
	Endorsements []Member
	IsOpen       bool
}
