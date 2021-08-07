package bounty

import (
	"encoding/json"
	"hackpoints/datastore"
	"hackpoints/models"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/shaj13/go-guardian/v2/auth"
)

type BountyServer struct {
	Store datastore.BountyStore
}

func validateNewBounty(b models.Bounty) error {
	if len(b.Title) == 0 {
		return ErrNoTitle
	}
	if len(b.Description) < 15 {
		return ErrBadDescription
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

	err = b.Store.NewBounty(*bounty)
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

	err = b.Store.UpdateBounty(*bounty)
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
		log.Debug("could not decode a bounty", err)
	}

	bounties, err := b.Store.GetBounties(*bounty)
	if err != nil {
		http.Error(w, ErrBountyNotFound.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	j, _ := json.Marshal(bounties)
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

	err = b.Store.EndorseBounty(*bounty, models.Member{Email: user.GetUserName()})
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

	err = b.Store.UpdateBounty(*bounty)
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
