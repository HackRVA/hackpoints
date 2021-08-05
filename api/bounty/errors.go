package bounty

import "errors"

var (
	ErrNoTitle                = errors.New("bounty needs a title")
	ErrBadDescription         = errors.New("bounty needs a longer description")
	ErrCantCreateClosedBounty = errors.New("bounty can't be created closed")
	ErrUpdatingBounty         = errors.New("bounty could not be updated")
	ErrBountyNotFound         = errors.New("bounty not found")
)
