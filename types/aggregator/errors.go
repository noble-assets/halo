// Copyright 2024 NASD Inc.
//
// Use of this source code is governed by a BSL-style
// license that can be found in the LICENSE file or at
// https://mariadb.com/bsl11.

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
