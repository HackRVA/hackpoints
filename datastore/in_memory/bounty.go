package in_memory

import (
	"errors"
	"hackpoints/models"
	"strconv"
)

var bounties = map[string]models.Bounty{}

type Store struct{}

func (b *Store) NewBounty(m models.Bounty) error {
	m.ID = strconv.Itoa(len(bounties) + 1)
	m.IsOpen = true
	bounties[m.ID] = m
	return nil
}

func (b *Store) UpdateBounty(m models.Bounty) error {
	if _, ok := bounties[m.ID]; !ok {
		return errors.New("can't update a bounty doesn't exist")
	}
	bounties[m.ID] = m
	return nil
}

func (b *Store) GetBounties(m models.Bounty) ([]models.Bounty, error) {
	if m.ID == "" {
		return bountyMapToSlice(bounties), nil
	}
	if _, ok := bounties[m.ID]; !ok {
		return []models.Bounty{}, errors.New("can't update a bounty doesn't exist")
	}
	return []models.Bounty{bounties[m.ID]}, nil
}

func bountyMapToSlice(b map[string]models.Bounty) []models.Bounty {
	var slice []models.Bounty

	for _, v := range b {
		slice = append(slice, v)
	}

	return slice
}

func (b *Store) EndorseBounty(bty models.Bounty, m models.Member) error {
	if _, ok := bounties[bty.ID]; !ok {
		return errors.New("could not find the bounty you were looking for")
	}

	if !bounties[bty.ID].IsOpen {
		return errors.New("bounty is closed and can no longer receive endorsements")
	}

	for _, k := range bounties[bty.ID].Endorsements {
		if k.ID == m.ID {
			return errors.New("you have already endorsed this bounty")
		}
	}

	endorsements := []models.Member{m}
	endorsements = append(endorsements, bounties[bty.ID].Endorsements...)

	bounties[bty.ID] = models.Bounty{
		ID:           bounties[bty.ID].ID,
		Title:        bounties[bty.ID].Title,
		Description:  bounties[bty.ID].Description,
		Endorsements: endorsements,
		IsOpen:       bounties[bty.ID].IsOpen,
	}

	return nil
}

func ClearBounties() {
	bounties = make(map[string]models.Bounty)
}
