// Copyright 2024 NASD Inc.
//
// Use of this source code is governed by a BSL-style
// license that can be found in the LICENSE file or at
// https://mariadb.com/bsl11.

package aggregator

import "cosmossdk.io/errors"

var (
	Codespace = "halo/aggregator"

	ErrNoReporter      = errors.Register(Codespace, 1, "there is no reporter")
	ErrInvalidReporter = errors.Register(Codespace, 2, "signer is not reporter")
	ErrSameReporter    = errors.Register(Codespace, 3, "provided reporter is the current reporter")

	ErrAlreadyReported  = errors.Register(Codespace, 4, "round already reported")
	ErrInvalidNextPrice = errors.Register(Codespace, 5, "next price is invalid")
)
