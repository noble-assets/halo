// Copyright 2024 NASD Inc.
//
// Use of this source code is governed by a BSL-style
// license that can be found in the LICENSE file or at
// https://mariadb.com/bsl11.

package aggregator

import "encoding/binary"

const SubmoduleName = "halo-aggregator"

var (
	OwnerKey       = []byte("aggregator/owner")
	LastRoundIDKey = []byte("aggregator/last_round_id")
	NextPriceKey   = []byte("aggregator/next_price")
	RoundPrefix    = []byte("aggregator/round/")
)

func RoundKey(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return append(RoundPrefix, bz...)
}
