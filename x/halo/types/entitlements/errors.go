package entitlements

import "cosmossdk.io/errors"

var (
	Codespace = "halo/entitlements"

	ErrNoOwner      = errors.Register(Codespace, 1, "there is no owner")
	ErrInvalidOwner = errors.Register(Codespace, 2, "signer is not owner")
	ErrSameOwner    = errors.Register(Codespace, 3, "provided owner is the current owner")
	ErrUnauthorized = errors.Register(Codespace, 4, "unauthorized")
)
