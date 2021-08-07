package in_memory

import (
	"hackpoints/models"
)

func (i *Store) GetScore() (int, error) {
	count := 0
	bounties, err := i.GetBounties(models.Bounty{})
	if err != nil {
		return count, err
	}
	for _, v := range bounties {
		count = count + len(v.Endorsements)
	}
	return count, nil
}
