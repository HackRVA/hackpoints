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
	Body BountyID
}

// swagger:parameters bountyCloseRequest
type bountyCloseRequest struct {
	// in: body
	Body BountyID
}

// swagger:parameters bountyGetRequest
type bountyGetRequest struct {
	// in: body
	Body BountyID
}

// swagger:response bountyResponse
type bountyResponse struct {
	Body Bounty
}

type BountyID struct {
	ID string `json:"id"`
}

type Bounty struct {
	ID           string   `json:"id"`
	Title        string   `json:"title"`
	Description  string   `json:"description"`
	Endorsements []Member `json:"endorsements"`
	IsOpen       bool     `json:"isOpen"`
}
