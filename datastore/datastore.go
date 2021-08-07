package datastore

import "hackpoints/models"

type BountyStore interface {
	NewBounty(models.Bounty) error
	UpdateBounty(models.Bounty) error
	GetBounties(models.Bounty) ([]models.Bounty, error)
	EndorseBounty(models.Bounty, models.Member) error
}

type ScoreStore interface {
	GetScore() (int, error)
}

type UserStore interface {
	GetMemberByEmail(email string) (models.Member, error)
	RegisterUser(models.Credentials) error
}
