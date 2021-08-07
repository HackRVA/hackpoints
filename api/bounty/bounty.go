package bounty

import (
	"encoding/json"
	"hackpoints/models"
	"net/http"

	"github.com/shaj13/go-guardian/v2/auth"
)

type BountyServer struct {
	Store models.BountyStore
}

func validateNewBounty(b models.Bounty) error {
	if len(b.Title) == 0 {
		return ErrNoTitle
	}
	if len(b.Description) < 15 {
		return ErrBadDescription
	}
	if !b.IsOpen {
		return ErrCantCreateClosedBounty
	}
	return nil
}

func (b *BountyServer) New(w http.ResponseWriter, r *http.Request) {
	bounty := &models.Bounty{}
	err := json.NewDecoder(r.Body).Decode(bounty)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = validateNewBounty(*bounty)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = b.Store.New(*bounty)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	j, _ := json.Marshal(models.EndpointSuccess{
		Ack: true,
	})
	w.Write(j)
}

func (b *BountyServer) Update(w http.ResponseWriter, r *http.Request) {
	bounty := &models.Bounty{}
	err := json.NewDecoder(r.Body).Decode(bounty)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = b.Store.Update(*bounty)
	if err != nil {
		http.Error(w, ErrUpdatingBounty.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	j, _ := json.Marshal(models.EndpointSuccess{
		Ack: true,
	})
	w.Write(j)
}

func (b *BountyServer) Get(w http.ResponseWriter, r *http.Request) {
	bounty := &models.Bounty{}
	err := json.NewDecoder(r.Body).Decode(bounty)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response, err := b.Store.Get(*bounty)
	if err != nil {
		http.Error(w, ErrBountyNotFound.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	j, _ := json.Marshal(response)
	w.Write(j)
}

func (b *BountyServer) Endorse(w http.ResponseWriter, r *http.Request) {
	bounty := &models.Bounty{}
	err := json.NewDecoder(r.Body).Decode(bounty)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user := auth.UserFromCtx(r.Context())

	err = b.Store.Endorse(*bounty, models.Member{Email: user.GetUserName()})
	if err != nil {
		http.Error(w, ErrUpdatingBounty.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	j, _ := json.Marshal(models.EndpointSuccess{
		Ack: true,
	})
	w.Write(j)
}

func (b *BountyServer) Close(w http.ResponseWriter, r *http.Request) {
	bounty := &models.Bounty{}
	err := json.NewDecoder(r.Body).Decode(bounty)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	bounty.IsOpen = false

	err = b.Store.Update(*bounty)
	if err != nil {
		http.Error(w, ErrUpdatingBounty.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	j, _ := json.Marshal(models.EndpointSuccess{
		Ack: true,
	})
	w.Write(j)
}
