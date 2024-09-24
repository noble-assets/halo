package entitlements

import "encoding/binary"

const SubmoduleName = "halo-entitlements"

var (
	OwnerKey         = []byte("entitlements/owner")
	PausedKey        = []byte("entitlements/paused")
	PublicPrefix     = []byte("entitlements/public/")
	CapabilityPrefix = []byte("entitlements/capability/")
	UserPrefix       = []byte("entitlements/user/")
)

func PublicKey(method string) []byte {
	return append(PublicPrefix, []byte(method)...)
}

func CapabilityKey(method string) []byte {
	return append(CapabilityPrefix, []byte(method)...)
}

func CapabilityRoleKey(method string, role Role) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, uint64(role))
	return append(CapabilityKey(method), bz...)
}

func UserKey(address []byte) []byte {
	return append(UserPrefix, address...)
}

func UserRoleKey(address []byte, role Role) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, uint64(role))
	return append(UserKey(address), bz...)
}
