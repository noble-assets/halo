// Copyright 2024 NASD Inc.
//
// Use of this source code is governed by a BSL-style
// license that can be found in the LICENSE file or at
// https://mariadb.com/bsl11.

package entitlements

import "encoding/binary"

var (
	OwnerKey         = []byte("entitlements/owner")
	PausedKey        = []byte("entitlements/paused")
	PublicPrefix     = []byte("entitlements/public/")
	CapabilityPrefix = []byte("entitlements/capability/")
	UserPrefix       = []byte("entitlements/user/")
)

func CapabilityRoleKey(method string, role Role) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, uint64(role))
	return append([]byte(method), bz...)
}
