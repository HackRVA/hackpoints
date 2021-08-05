package in_memory

import (
	"errors"
	"hackpoints/models"
)

var bounties = map[string]models.Bounty{}

type InMemoryBountyStore struct{}

func (b *InMemoryBountyStore) New(m models.Bounty) error {
	bounties[m.ID] = m
	return nil
}

func (b *InMemoryBountyStore) Update(m models.Bounty) error {
	if _, ok := bounties[m.ID]; !ok {
		return errors.New("can't update a bounty doesn't exist")
	}
	bounties[m.ID] = m
	return nil
}

func (b *InMemoryBountyStore) Get(m models.Bounty) (models.Bounty, error) {
	if _, ok := bounties[m.ID]; !ok {
		return models.Bounty{}, errors.New("can't update a bounty doesn't exist")
	}
	return bounties[m.ID], nil
}

func (b *InMemoryBountyStore) GetAll() []models.Bounty {
	return []models.Bounty{}
}

func (b *InMemoryBountyStore) Endorse(models.Bounty) error {
	return nil
}
