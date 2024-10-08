// Copyright 2024 NASD Inc.
//
// Use of this source code is governed by a BSL-style
// license that can be found in the LICENSE file or at
// https://mariadb.com/bsl11.

package types

import "cosmossdk.io/errors"

var (
	ErrNoOwner                  = errors.Register(ModuleName, 1, "there is no owner")
	ErrInvalidOwner             = errors.Register(ModuleName, 2, "signer is not owner")
	ErrSameOwner                = errors.Register(ModuleName, 3, "provided owner is the current owner")
	ErrInvalidFundAdmin         = errors.Register(ModuleName, 4, "signer is not a fund admin")
	ErrInvalidLiquidityProvider = errors.Register(ModuleName, 5, "signer is not a liquidity provider")
	ErrInvalidSignature         = errors.Register(ModuleName, 6, "invalid withdrawal signature")
)
