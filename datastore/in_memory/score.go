package in_memory

import "hackpoints/models"

type InMemoryScoreStore struct {
	BountyStore models.BountyStore
}

func (i *InMemoryScoreStore) Get() (int, error) {
	count := 0
	bounties, err := i.BountyStore.Get(models.Bounty{})
	if err != nil {
		return count, err
	}
	for _, v := range bounties {
		count = count + len(v.Endorsements)
	}
	return count, nil
}
