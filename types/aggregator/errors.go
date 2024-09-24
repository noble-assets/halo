package aggregator

import "cosmossdk.io/errors"

var (
	Codespace = "halo/aggregator"

	ErrNoOwner      = errors.Register(Codespace, 1, "there is no owner")
	ErrInvalidOwner = errors.Register(Codespace, 2, "signer is not owner")
	ErrSameOwner    = errors.Register(Codespace, 3, "provided owner is the current owner")

	ErrAlreadyReported  = errors.Register(Codespace, 4, "round already reported")
	ErrInvalidNextPrice = errors.Register(Codespace, 5, "next price is invalid")
)
