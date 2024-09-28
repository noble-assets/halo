// Copyright 2024 NASD Inc.
//
// Use of this source code is governed by a BSL-style
// license that can be found in the LICENSE file or at
// https://mariadb.com/bsl11.

package utils

import "cosmossdk.io/math"

func MustParseInt(s string) math.Int {
	res, _ := math.NewIntFromString(s)
	return res
}
